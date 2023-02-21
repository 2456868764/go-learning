package engine

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Context struct {

	W http.ResponseWriter
	R *http.Request
	Path string
	Method string
	// Keys map 读写锁
	mu sync.RWMutex
	// 每个请求独立上下文传值用
	Keys map[string]any
	//路由匹配数据
	PathParams map[string]string
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	context := &Context{
		W: w,
		R: r,
		PathParams: make(map[string]string),
	}
	if r!= nil {
		context.Method = r.Method
		context.Path = r.URL.Path
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





// Set is used to store a new key/value pair exclusively for this context.
// It also lazy initializes  c.Keys if it was not used previously.
func (c *Context) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.Keys == nil {
		c.Keys = make(map[string]any)
	}
	c.Keys[key] = value
}

// Get returns the value for the given key, ie: (value, true).
// If the value does not exist it returns (nil, false)
func (c *Context) Get(key string) (value any, exists bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists = c.Keys[key]
	return
}

// GetString returns the value associated with the key as a string.
func (c *Context) GetString(key string) (s string) {
	if val, ok := c.Get(key); ok && val != nil {
		s, _ = val.(string)
	}
	return
}

// GetInt returns the value associated with the key as an integer.
func (c *Context) GetInt(key string) (i int) {
	if val, ok := c.Get(key); ok && val != nil {
		i, _ = val.(int)
	}
	return
}






