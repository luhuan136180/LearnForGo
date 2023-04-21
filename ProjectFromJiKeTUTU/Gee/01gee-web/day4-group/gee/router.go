package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node //存储路由树的根节点，key是method类型
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

//用于获取分割好的路由，用于在路由树中查找
func parsePattern(pattern string) []string {
	//按照“/"切割
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

//传入这里的pattern是完整的路由路径，
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern) //将其切割

	key := method + "-" + pattern //将完整路径结合method 组合生成key 用于存储对应的路由函数
	_, ok := r.roots[method]      //查询该路由的method对应根节点是否存在
	if !ok {                      //不存在
		r.roots[method] = &node{} //懒汉式，没有的话才初始化
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler //注册路由和其方法的映射
}

//返回一个拥有匹配成功的完整路由路径的节点和存储动态路径参数的一个切片，参数数量和类容不确定
func (r *router) getRoute(method, path string) (*node, map[string]string) {
	searchParts := parsePattern(path) //请求路由的解析切片
	params := make(map[string]string)
	root, ok := r.roots[method] //先查询此method的路由树是否存在
	if !ok {                    //不存在，不可能有对应路由，返回nil
		return nil, nil
	}
	n := root.search(searchParts, 0) //将请求的路由路径转化的切片传入，用于在路由树中查询
	if n != nil {                    //查询成功
		parts := parsePattern(n.pattern) //将真实的匹配成功的路由路径转换一下，用于映射路由函数等
		for index, part := range parts {
			if part[0] == ':' {
				//这时需要将请求路由路径对应的位置作为参数传入参数组合中
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) getRoutes(method string) []*node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.travel(&nodes)
	return nodes
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
