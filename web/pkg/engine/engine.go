package engine

import (
	"net/http"
)


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
		//router: NewMapBasedRouter(),
		router: NewTreeBasedRouter(),
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
