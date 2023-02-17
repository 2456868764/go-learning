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
  build            Build binary with the crane manager.
  image            Build docker image with the crane manager.
  push-image       Push images.

```

## 具体实现

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




