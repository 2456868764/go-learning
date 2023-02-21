package engine

import "net/http"

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

