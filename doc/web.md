# web实战项目需求和目标
学习和工作中用过比较多Web框架，对Web框架的原理需要进一步了解比如怎么注册路由，如何路由匹配，如何灵活运用 Web 框架提供的middleware方案解决登录校验、鉴权、日志、tracing、logging,限流等问题
项目实践练习设计一个Web Server，该 Web Server 框架这个教程的很多设计，包括源码，参考了Gin,可以看到很多Gin的影子。

## 项目目标
1. 学习下面主要知识点
* http.Handler(Engine)
* 上下文设计(Context)
* 前缀路由树Trie树路由(Router)，支持路由通配符匹配、路径参数
* 分组控制(Group)
* 中间件(Middleware)
* HTML模板(Template)
* 错误恢复(Panic Recover)
* 优雅关闭
* 统一错误处理

2. 实现httpbin大部分功能  http://www.httpbin.org/

## 项目打包准备
用 docker 和 make 构建自动化工具，Docker和 Make 自动构建介绍参考 https://github.com/2456868764/k8s-learning
1. Dockerfile

```shell
# Build the manager binary
FROM golang:1.19.4-alpine3.16 as builder
ARG LDFLAGS
ARG PKGNAME
ARG BUILD
ENV GO111MODULE=on \
    CGO_ENABLED=0
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.mod
#COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN if [[ "${BUILD}" != "CI" ]]; then go env -w GOPROXY=https://goproxy.io,direct; fi
RUN go env
RUN go mod download

# Copy the go source
COPY api api/
COPY pkg pkg/
COPY cmd cmd/

# Build
RUN env
RUN go build -ldflags="${LDFLAGS}" -a -o httpbin cmd/main.go

FROM alpine:3.15.3
WORKDIR /app
ARG PKGNAME
COPY --from=builder /app/httpbin .
CMD ["./httpbin"]
```
2. Makefile
```shell
make help

Usage:
  make <target>

General
  help             Display this help.
  fmt              Run go fmt against code.
  vet              Run go vet against code.
  lint             Run golang lint against code
  test             Run tests.
  build            Build binary with the httpbin.
  image            Build docker image with the httpbin.
  push-image       Push images.

```

## 具体实现

### 第一步：设计Engine、简单路由注册、项目结构

1. 要实现 Handler接口
```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

2. 设计一个Engine结构体实现这个接口，同时定义一个handler函数 HandlerFunc，接收 http.ResponseWriter,  *http.Request参数

```go

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}


```
```go
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	routeKey := r.Method + "-" + r.URL.Path
	handler, ok := e.router[routeKey]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Found Method: %s Path: %s", r.Method, r.URL.Path)
		return
	}
	handler(w, r)
}
```
3. 同时实现路由注册
```go

func (e *Engine) AddRoute(method string, pattern string, handler HandlerFunc) {
	routeKey := method + "-" + pattern
	e.router[routeKey] = handler
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.AddRoute(http.MethodGet, pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.AddRoute(http.MethodPost, pattern, handler)
}

```
4. 同时启动服务

```go
func (e *Engine) Run(addr string) error {
	err := http.ListenAndServe(addr, e)
	return err
}

```

5. 一个简单web服务就起来拉

```go
func main() {
	engine := engine.New()
	engine.GET("/headers", v1.GetHeaders)
	engine.GET("/ip", v1.GetIP)
	engine.GET("/user-agent", v1.GetUserAgent)
	engine.Run(":8080")
}

```
6. 项目结构

```shell
.
├── Dockerfile
├── Makefile
├── api
│   └── v1
│       └── api.go
├── bin
│   └── httpbin
├── cmd
│   └── main.go
├── go.mod
└── pkg
    └── engine
        └── engine.go


```

### 第二步：添加Context请求上下文和定义Routable & Router接口
1. Context, 请求上下文。 我们为什么要构建一个请求上下文

* 封装 net/http下ServeHTTP(w http.ResponseWriter, r *http.Request)接口中两个参数
* 每个请求独立上下文，可以做很多事情，比如传递上下相关参数，便于后面开发Filter chain
* 封装 公共工具方法，方便使用，减少重复代码

```go

type Context struct {

	W http.ResponseWriter
	R *http.Request
	Path string
	Method string
	// Keys map 读写锁
	mu sync.RWMutex
	// 每个请求独立上下文传值用
	Keys map[string]any
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	context := &Context{
		W: w,
		R: r,
		Path: r.URL.Path,
		Method: r.Method,
	}
	return context
}

func (c *Context) ReadJsonObject(object any) error {
	body, err := io.ReadAll(c.R.Body)
	if err!= nil {
		return err
	}

	return json.Unmarshal(body, object)
}

func (c *Context) ReponseJson(httpStatus int, object any) error {
	c.SetHeader("Content-Type", "application/json")
	c.Status(httpStatus)
	bytes, err := json.Marshal(object)
	if err!= nil {
		return err
	}
	_, err = c.W.Write(bytes)
	if err!= nil {
		return err
	}

	return nil
}

func (c *Context)OKJson(object any) error {
	return c.ReponseJson(http.StatusOK, object)
}

func (c *Context)ServerErrorJson(object any) error {
	return c.ReponseJson(http.StatusInternalServerError, object)
}

func (c *Context)BadRequestJson(object any) error {
	return c.ReponseJson(http.StatusBadRequest, object)
}


func (c *Context) PostForm(key string) string {
	return c.R.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.R.URL.Query().Get(key)
}

func (c *Context) GetHeader(key string) string {
	return c.R.Header.Get(key)
}

func (c *Context) Status(code int) {
	c.W.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.W.Header().Set(key, value)
}

func (c *Context) StringFormat(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.W.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) StringOk(body string) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(http.StatusOK)
	io.WriteString(c.W, body)
}

```

2. 定义Routable & Router接口

定义Routable & Router接口目的
* 把 Engine 里 router 独立抽象出来一个接口路由器，Engine 把添加路由，处理路由功能交个这个接口具体实现类，这样可以解耦
* 同时提供Router接口不同实现，目前实现基于Map结构，后续会实现基于前缀树更强功能路由器

```go

// Router 定义路由接口，可以用不同的实现，可以基于 map 和 前缀树的实现
type Router interface {
	ServerHTTP(c * Context)
	Routable
}

func NewMapBasedRouter() Router {
	router := &MapBasedRouter{
		handlers: make(map[string]HandlerFunc),
	}
	return router
}

type MapBasedRouter struct {
	handlers map[string]HandlerFunc
}

func (m *MapBasedRouter) ServerHTTP(c *Context) {
	routeKey := c.Method + "-" + c.Path
	handler, ok := m.handlers[routeKey]
	if !ok {
		c.StringFormat(http.StatusNotFound, "Not Found Method: %s Path: %s", c.Method, c.Path)
		return
	}
	handler(c)
}

func (m *MapBasedRouter) AddRoute(method string, pattern string, handlerFunc HandlerFunc) error {
	routeKey := method + "-" + pattern
	m.handlers[routeKey] = handlerFunc
	return nil
}
```
3. 调整Engine实现

```go

// Routable 可以路由
type Routable interface {
	// AddRoute 添加一个路由，命中该路由的调用 handlerFunc 代码
	AddRoute(method string, pattern string, handlerFunc HandlerFunc) error
}

// HandlerFunc 某个路由对应具体执行
type HandlerFunc func(c *Context)

type Engine struct {
	router Router
}

func New() *Engine {
	engine := &Engine{
		router: NewMapBasedRouter(),
	}
	return engine
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)
	e.router.ServerHTTP(c)
}

func (e *Engine) AddRoute(method string, pattern string, handler HandlerFunc) error {
	e.router.AddRoute(method, pattern, handler)
	return nil
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.AddRoute(http.MethodGet, pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.AddRoute(http.MethodPost, pattern, handler)
}

func (e *Engine) Run(addr string) error {
	err := http.ListenAndServe(addr, e)
	return err
}
```
Engine实际就非常简单，都是委托给Router

4. 调整api 实现

```go

func GetUserAgent(c *engine.Context) {
	ip := c.GetHeader("User-Agent")
	c.StringOk(fmt.Sprintf("User-Agent=%s\n", ip))
}

func GetIP(c *engine.Context) {
	ip := c.GetHeader("REMOTE-ADDR")
	c.StringOk(fmt.Sprintf("IP=%s\n", ip))
}

func GetHeaders(c *engine.Context) {
	content := ""
	for k, v := range c.R.Header {
		content += fmt.Sprintf("%s=%s\n", k, v)
	}
	c.StringOk(content)
}

```

