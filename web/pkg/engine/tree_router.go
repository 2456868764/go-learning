package engine

import (
	"errors"
	"net/http"
	"strings"
)

var ErrorInvalidRouterPathPattern = errors.New("invalid router path pattern")
var ErrorInvalidRouterMethod = errors.New("invalid http method")

var supportedMethods = [5]string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodDelete,
	http.MethodPatch,
}

type TreeBasedRouter struct {
	// 每个支持方法相对应一个前缀树
	routeForest map[string]*node
}

func NewTreeBasedRouter() Router {
	forest := make(map[string]*node, len(supportedMethods))
	for _, method := range supportedMethods {
		forest[method] = newNodeRoot(method)
	}
	router := &TreeBasedRouter{
		routeForest: forest,
	}
	return router
}

func (t *TreeBasedRouter) ServerHTTP(c *Context) {
	handler, ok := t.findRoute(c.Method, c.Path, c)
	if !ok {
		c.StringFormat(http.StatusNotFound, "Not Found Method: %s Path: %s", c.Method, c.Path)
		return
	}
	handler(c)
}

func (t *TreeBasedRouter) AddRoute(method string, pattern string, handler HandlerFunc) error {
	err := validRoutePathPattern(pattern)
	if err != nil {
		return err
	}

	// 根据请求方法类型找到对应的前缀树根节点
	rootNode, ok := t.routeForest[method]
	if !ok {
		return ErrorInvalidRouterMethod
	}

	// 把路由分割成数组， 比如/order/detail, 分割成【order, detail]
	paths := strings.Split(strings.Trim(pattern, "/"), "/")

	currNode := rootNode
	for index, path := range paths {
		child, found := currNode.findChild(path, nil)
		if found && child.nodeType != nodeTypeAny {
			// 找到，继续找
			currNode = child
		} else {
			// 没有找到，后面的路由作为当前节点子节点添加，添加完成，返回叶节点，跳出 for 循环
			currNode = currNode.addChild(paths[index:], handler)
			break
		}
	}

	// 到这里, 重新设置节点handler 和 路由节点标志
	currNode.handler = handler
	currNode.end = true

	return nil
}

func validRoutePathPattern(pattern string) error {
	// 目前只接受 /* 这个路由风格， * 必须是最后一个字符，同时前一个字符是 /
	starPos := strings.Index(pattern, "*")
	if starPos > 0 {
		if starPos != len(pattern)-1 {
			return ErrorInvalidRouterPathPattern
		}
		if pattern[starPos-1] != '/' {
			return ErrorInvalidRouterPathPattern
		}
	}

	return nil
}

func (t *TreeBasedRouter) findRoute(method string, path string, c *Context) (HandlerFunc, bool) {
	paths := strings.Split(strings.Trim(path, "/"), "/")

	// 方法不支持
	rootNode, ok := t.routeForest[method]
	if !ok {
		return nil, false
	}

	currNode := rootNode
	for _, subPath := range paths {
		child, found := currNode.findChild(subPath, c)
		if !found {
			return nil, false
		}
		currNode = child
	}

	// 找到了，判断一下是否注册的路由
	if !currNode.end {
		// 比如这种场景
		// 比如注册了 /order/detail/info
		// 但是访问 /order
		return nil, false
	}

	return currNode.handler, true
}
