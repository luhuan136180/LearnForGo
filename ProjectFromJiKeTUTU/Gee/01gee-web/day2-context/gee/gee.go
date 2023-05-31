package gee

import "net/http"

type HandlerFunc func(ctx *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

//实现ServeHTTP接口--
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req) //初始化一个自定义的上下文，包含了w和req，以及一些其他属性（本版本只添加了url和method，以后还可以写入动态参数，中间件等）
	engine.router.handle(c) //根据请求的路由路径选择路由方法--将具体找路由方法的操作封装到该函数中--
}
