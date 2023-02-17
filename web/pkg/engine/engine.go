package engine

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	engine := &Engine{
		router: make(map[string]HandlerFunc),
	}
	return engine
}

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

func (e *Engine) Run(addr string) error {
	err := http.ListenAndServe(addr, e)
	return err
}
