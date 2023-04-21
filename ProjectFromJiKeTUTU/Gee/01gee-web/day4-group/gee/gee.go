package gee

import (
	"log"
	"net/http"
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

//初始化gee引擎
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine} //初始化一个路由组（根组）--引擎本身
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

//定义组来创建一个新的RouterGroup
//所有组共享同一个Engine实例
//该函数创建一个新的group组？通过父路由组的调用，传入新路由组的共有前缀，引用相同的引擎，并将自己添加入引擎的路由组注册切片数组中
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix, //生成新的组的共有组前缀
		parent: group,                 //赋值父组信息
		engine: engine,                //赋值唯一引擎
	}
	engine.groups = append(engine.groups, newGroup) //在项目引擎的路由组表中添加生成的组的信息
	return newGroup                                 //huitui
}

//有关路由方法和路由路径的函数，可以交给routergroup实现了
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	//一下拼接prefix的方式可以将group的parent元素舍弃掉，
	pattern := group.prefix + comp //结合当前的路由组的路径加上组内的路径生成完整的url路径（不携带query参数版）
	log.Printf("Route %4s - %s", method, pattern)

	//总之我现在还是不是非常理解
	//注：Engine从某种意义上继承了RouterGroup的所有属性和方法，因为 (*Engine).engine 是指向自己的。
	//这样实现，我们既可以像原来一样添加路由，也可以通过分组添加路由。——————》理由是 engine继承了RouterGroup，RouterGroup的方法，可以被engine直接调用
	group.engine.router.addRoute(method, pattern, handler) //实现路由映射
}

//
// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

//一下基于引擎的方法
// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
