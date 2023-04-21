package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc定义了gee使用的请求处理程序
type HandlerFunc func(w http.ResponseWriter, req *http.Request)

//gee引擎实例 实现了Handler接口
type Engine struct {
	//key：请求方法和静态路由地址
	router map[string]HandlerFunc //用于存储路由和对应的请求处理函数——路由映射表
}

//实现了Handler接口的ServeHTTP方法
//ServeHTTP方法实现了Handler接口的ServeHTTP方法，用于处理HTTP请求。具体来说，当收到一个HTTP请求时，它会将请求方法和URL作为map的key，
//查找对应的路由处理函数，如果找到则调用该函数处理请求，如果找不到则返回404错误。
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//解析请求的路径，查找路由映射表，如果查到就这行注册的处理方法，如果查不到，就返回404 not found

	key := req.Method + "-" + req.URL.Path     //key：请求方法和静态路由地址
	if handler, ok := engine.router[key]; ok { //寻找是否有这个对应映射
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND : %s\n", req.URL)
	}
}

//构建一个引擎的实例
func New() *Engine { //初始化引擎实例
	return &Engine{
		router: make(map[string]HandlerFunc), //初始化一个新的路由表,此时为空
	}
}

//注册路由函数，将路由相应函数和其想对应的url存入引擎的路由表中，
//key：请求方式（get，post等） - url
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET定义添加GET请求的方法  例：GET("/login",server.Login)
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run定义启动http服务器的方法  对http.ListenAndServe（）地包装
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine) //输入监听端口，和使用的引擎
}
