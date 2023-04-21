package gee

import (
	"log"
	"net/http"
	"strings"
)

// 声明类型：HandlerFunc defines the request handler used by gee
type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
type (
	//总之就是很高级思路，不是很理解，
	//
	RouterGroup struct {
		prefix      string        //该路由组的共有前缀
		middlewares []HandlerFunc // 支持中间件
		parent      *RouterGroup  // 支持嵌套--当前gee中使用了prefix嵌套，使得该元素可以舍弃
		engine      *Engine       //所有组共享一个Engine实例,此属性每一个项目唯一
	}

	Engine struct {
		*RouterGroup                //嵌套，类似继承
		router       *router        //引擎锁需要的路由信息，包括路由方法映射表和路由节点树
		groups       []*RouterGroup // 储存所有路由组组
	}
)

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

//中间件的添加基于路由组结构体
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)

}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc //初始化一个中间件执行列表--切片
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			//只要url能匹配上该组的prefix，则将这个路由组绑定的中间件全部添加在这个路由函数需要执行的中间件链表中
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)  //初始化上下文
	c.handlers = middlewares //将该路由需要的中间件赋值入上下文的中间件
	engine.router.handle(c)
}
