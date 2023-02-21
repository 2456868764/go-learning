package engine

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)


func TestTreeBasedRouter_AddRoute(t *testing.T) {
	handler := NewTreeBasedRouter().(*TreeBasedRouter)
	assert.Equal(t, len(supportedMethods), len(handler.routeForest))

	postRootNode := handler.routeForest[http.MethodPost]
	err := handler.AddRoute(http.MethodPost, "/blog", func(c *Context) {})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(postRootNode.children))
	blogNode := postRootNode.children[0]
	assert.NotNil(t, blogNode)
	assert.Equal(t, "blog", blogNode.nodePathPattern)
	assert.NotNil(t, blogNode.handler)
	assert.Empty(t, blogNode.children)


	err = handler.AddRoute(http.MethodPost, "/blog/detail", func(c *Context) {})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(blogNode.children))
	detailNode := blogNode.children[0]
	assert.NotNil(t, detailNode)
	assert.Equal(t, "detail", detailNode.nodePathPattern)
	assert.NotNil(t, detailNode.handler)
	assert.Empty(t, detailNode.children)

	// 测试重复添加
	err = handler.AddRoute(http.MethodPost, "/blog", func(c *Context) {})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(postRootNode.children))
	blogNode = postRootNode.children[0]
	assert.NotNil(t, blogNode)
	assert.Equal(t, "blog", blogNode.nodePathPattern)
	assert.NotNil(t, blogNode.handler)
	assert.Equal(t, 1, len(blogNode.children))


	err = handler.AddRoute(http.MethodPost, "/blog/add", func(c *Context) {})
	assert.Nil(t, err)
	assert.Equal(t, 2, len(blogNode.children))
	addNode := blogNode.children[1]
	assert.NotNil(t, addNode)
	assert.Equal(t, "add", addNode.nodePathPattern)
	assert.NotNil(t, addNode.handler)
	assert.Empty(t, addNode.children)



	err = handler.AddRoute(http.MethodPost, "/user", func(c *Context) {})
	assert.Nil(t, err)
	assert.Equal(t, 2, len(postRootNode.children))
	userNode := postRootNode.children[1]
	assert.NotNil(t, userNode)
	assert.Equal(t, "user", userNode.nodePathPattern)
	assert.NotNil(t, userNode.handler)
	assert.Empty(t, userNode.children)

	err = handler.AddRoute(http.MethodPost, "/user/:id", func(c *Context) {})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(userNode.children))
	idNode := userNode.children[0]
	assert.NotNil(t, idNode)
	assert.Equal(t, ":id", idNode.nodePathPattern)
	assert.Equal(t, nodeTypeParam, idNode.nodeType)
	assert.NotNil(t, idNode.handler)
	assert.Empty(t, idNode.children)


	err = handler.AddRoute(http.MethodPost, "/user/*/info", func(c *Context) {})
	assert.Equal(t, ErrorInvalidRouterPathPattern, err)


	err = handler.AddRoute(http.MethodPost, "/user/*", func(c *Context) {})
	assert.Nil(t, err)
	assert.Equal(t, 2, len(userNode.children))
	anyNode := userNode.children[1]
	assert.NotNil(t, anyNode)
	assert.Equal(t, "*", anyNode.nodePathPattern)
	assert.Equal(t, nodeTypeAny, anyNode.nodeType)
	assert.NotNil(t, idNode.handler)
	assert.Empty(t, idNode.children)

}

func TestTreeBasedRouter_findRoute(t1 *testing.T) {
	handler := NewTreeBasedRouter().(*TreeBasedRouter)
	context := NewContext(nil, nil)
	//rootNode := handler.routeForest[http.MethodPost]
	_ = handler.AddRoute(http.MethodPost, "/blog", func(c *Context) {})
	//blogNode := rootNode.children[0]
	//fmt.Printf("blog node:%+v\n", blogNode)
	fn, found := handler.findRoute(http.MethodPost, "/blog", context)
	assert.True(t1, found)
	assert.NotNil(t1, fn)

	_ = handler.AddRoute(http.MethodPost, "/blog/detail", func(c *Context) {})
	//fmt.Printf("detail node:%+v\n", blogNode.children[0])
	fn, found = handler.findRoute(http.MethodPost, "/blog/detail", context)
	assert.True(t1, found)
	assert.NotNil(t1, fn)
	//
	_ = handler.AddRoute(http.MethodPost, "/blog/:blogId", func(c *Context) {})
	//fmt.Printf("blogid node:%+v\n", blogNode.children[1])
	fn, found = handler.findRoute(http.MethodPost, "/blog/3333", context)
	assert.True(t1, found)
	assert.NotNil(t1, fn)
	assert.Equal(t1, "3333", context.PathParams["blogId"])


	// 参数路由
	_ = handler.AddRoute(http.MethodPost, "/order/*", func(c *Context) {})
	fn, found = handler.findRoute(http.MethodPost, "/order/checkout", context)
	assert.True(t1, found)
	assert.NotNil(t1, fn)

}