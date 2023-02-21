package engine

import (
	"sort"
	"strings"
)

const (
	//跟节点
	nodeTypeRoot = iota
	// * 节点
	nodeTypeAny
	// :param 参数节点
	nodeTypeParam
	// [] 正则匹配节点
	nodeTypeReg
	// 完全匹配相等节点
	nodeTypeStatic
)

type nodeMatchFunc func(path string, c *Context) bool

type node struct {
	// 自节点列表
	children []*node
	// 节点匹配到处理函数
	handler HandlerFunc

	// 节点匹配函数
	nodeMatchFunc nodeMatchFunc

	// 当前节点对应的匹配模式
	nodePathPattern string
	// 节点类型
	nodeType int

	// 表示这个节点是路由路径一个node, 这个为true有handlerFunc, 为false 没有handlerFunc
	end bool
}

func (n *node) findChild(path string, c *Context) (*node, bool) {
	foundNodes := make([]*node, 0, 2)

	for _, child := range n.children {
		if child.nodeMatchFunc(path, c) {
			foundNodes = append(foundNodes, child)
		}
	}

	if len(foundNodes) == 0 {
		return nil, false
	}

	// 在找到多个匹配子节点时，根据子节点类型排序，
	sort.Slice(foundNodes, func(i int, j int) bool {
		return foundNodes[i].nodeType < foundNodes[j].nodeType
	})

	return foundNodes[len(foundNodes)-1], true

}

func (n *node) addChild(paths []string, handler HandlerFunc) *node {
	currNode := n
	for _, path := range paths {
		child := newNode(path)
		currNode.children = append(currNode.children, child)
		currNode = child
	}

	// 到这里, 设置节点handler 和 路由节点标志
	currNode.handler = handler
	currNode.end = true
	return currNode
}

func newNodeRoot(pattern string) *node {
	return &node{
		children:        make([]*node, 0, 1),
		nodeType:        nodeTypeRoot,
		nodePathPattern: pattern,
		nodeMatchFunc: func(path string, c *Context) bool {
			v := "shoudn't be called"
			panic(v)
		},
	}
}

func newNodeStatic(pattern string) *node {
	return &node{
		children:        make([]*node, 0, 1),
		nodeType:        nodeTypeStatic,
		end:             false,
		nodePathPattern: pattern,
		nodeMatchFunc: func(path string, c *Context) bool {
			//fmt.Printf("pattern=%s, path=%s\n", pattern, path)
			return pattern == path && path != "*"
		},
	}
}

func newNodeAny() *node {
	return &node{
		children:        make([]*node, 0, 1),
		nodeType:        nodeTypeAny,
		end:             false,
		nodePathPattern: "*",
		nodeMatchFunc: func(path string, c *Context) bool {
			//fmt.Printf("pattern=%s, path=%s\n", "*", path)
			return true
		},
	}
}

func newNodeParam(pattern string) *node {
	paramName := pattern[1:]
	return &node{
		children:        make([]*node, 0, 1),
		nodeType:        nodeTypeParam,
		end:             false,
		nodePathPattern: pattern,
		nodeMatchFunc: func(path string, c *Context) bool {
			//fmt.Printf("pattern=%s, path=%s, paramName=%s\n ", pattern, path, paramName)
			if c != nil {
				c.PathParams[paramName] = path
			}
			return path != "*"
		},
	}
}

func newNode(pattern string) *node {
	if pattern == "*" {
		return newNodeAny()
	}
	if strings.HasPrefix(pattern, ":") {
		return newNodeParam(pattern)
	}
	return newNodeStatic(pattern)
}
