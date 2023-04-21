package gee

import "net/http"

/*
将和路由相关的方法和结构提取了出来，放到了一个新的文件中router.go，
方便我们下一次对 router 的功能进行增强，例如提供动态路由的支持。

router 的 handle 方法作了一个细微的调整，即 handler 的参数，变成了 Context。
*/

type router struct {
	handlers map[string]HandlerFunc
}

//初始化路由
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

//向路由表中添加路由映射
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

//该方法用于找对应路由和路由方法：：根据c.path和c.method 结合作为路由表的key 查找项目服务端是否有对应的路由方法，有，则执行
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c) //执行--调用对应的路由函数
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
