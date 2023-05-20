package gee

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

/*
html/template标准库和text/template标准库都是Go语言内置模板引擎，主要功能是根据模板和数据生成文本输出。

text/template标准库可以解析普通文本模板，支持if、range、with等基本控制语句，而html/template标准库是在text/template标准库的基础上进行了扩展，
支持HTML、JS、CSS等标记语言的编写。

与text/template标准库相比，html/template标准库在解析和渲染过程中会对输入进行一定的转义，防止输出的文本包含恶意JavaScript代码等攻击代码。
因此，在对用户输入进行渲染时，建议使用html/template标准库以提高应用程序的安全性。
*/

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

		//用于渲染html
		//*template.Template:内置模板引擎的核心结构体，主要用于存储和管理解析后的模板信息。
		//结构体中存储了解析后的模板结构树，以及渲染时需要使用的所有变量、方法等信息，可以通过该结构体提供的方法进行模板渲染。
		//先通过text/template或html/template标准库的Parse方法或ParseFiles方法解析模板得到*template.Template结构体，然后就可以调用该结构体提供的Execute方法对模板进行渲染。
		//其次，支持模板继承。在模板继承中，子模板会复用父模板中的一部分内容，并在此基础上扩展自己的内容。通过在template.Template结构体中指定父模板，就可以实现模板继承的功能。
		htmlTemplates *template.Template //Template类型是text/template包的Template类型的特化版本，用于生成安全的HTML文本片段。
		funcMap       template.FuncMap   //FuncMap类型定义了函数名字符串到函数的映射，自定义函数映射类型，**本质是一个map**** key是函数的名称，value是对应的函数。
	}
	/*示例：*template.Template和template.FuncMap的联合使用
	// 定义一个名为"upper"的函数，将指定字符串转换为大写字母
	func upper(s string) string {
	    return strings.ToUpper(s)
	}

	// 将"upper"函数映射到一个FuncMap变量中
	funcMap := template.FuncMap{
	    "upper": upper,
	}

	// 创建一个*template.Template结构体
	t := template.New("test").Funcs(funcMap)

	// 解析模板
	tmpl, _ := t .Parse("{{upper .}}")

	// 渲染模板，输出结果为"HELLO WORLD"
	tmpl.Execute(os.Stdout, "hello world")
	*/

)

// Default use Logger() & Recovery middlewares
func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

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

//添加关于服务端模板渲染的代码

//参数2:  http.FileSystem 是 Go 语言中定义文件系统访问的接口。
//静态文件服务中，我们通常使用 http.Dir 来实现 http.FileSystem 接口，可以使用绝对路径或者相对路径作为参数创建一个文件系统，供后续文件访问操作使用。
//参数1:relativePath:路由路径
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	//Join函数可以将任意数量的路径元素放入一个单一路径里，会根据需要添加斜杠。
	//absolutePath = group.prefix/relativePath
	absolutePath := path.Join(group.prefix, relativePath)

	//http.FileServer(fs)::接收一个实现了 http.FileSystem 接口的参数 fs;;创建一个可以处理静态文件的 http.Handler
	//func StripPrefix(prefix string, h Handler) Handler
	//StripPrefix::返回一个处理器，该处理器会将请求的URL.Path字段中给定前缀prefix去除后再交由h处理。StripPrefix会向URL.Path字段中没有给定前缀的请求回复404 page not found
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))

	return func(c *Context) {
		file := c.Param("filepath")

		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

//暴露给用户的方法::将磁盘上的某个文件夹root映射到路由relativePath
//relativePath:路由路径；；root:磁盘中的文件路径
//r.Static("/assets", "/usr/geektutu/blog/static")
//	---->
//访问localhost:9999/assets/js/geektutu.js，最终返回/usr/geektutu/blog/static/js/geektutu.js
//
func (group *RouterGroup) Static(relativePath string, root string) {
	//返回一个读取执行路径的文件（root）的路由方法，
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	//在传入的路由路径后面添加上用于指示是读文件的表示/*filepath，这样传给路由映射创建的函数
	urlPattern := path.Join(relativePath, "/*filepath")

	//注册路由方法，
	group.GET(urlPattern, handler)
}

//

// for custom render function设置自定义渲染函数
func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	//传参是一个已经写好的用于模板渲染的函数映射表，将其赋值到引擎的内置属性中，方便路由函数使用
	engine.funcMap = funcMap
}

//设置自定义加载模板
func (engine *Engine) LoadHTMLGlob(pattern string) {
	//template.New(""):New用给定的名称分配一个新的HTML模板。初始化创建了一个新的 *template.Template 对象，具有默认设置，并且没有任何模板信息。
	//通过该对象的 Parse 或 ParseFiles 方法来加载和解析模板文件
	/*示例:
	// 创建一个名为 "hello" 的新模板对象
	t := template.New("hello")

	// 加载并解析包含模板信息的字符串
	t.Parse("hello {{.}}")

	// 渲染模板，并将结果输出到控制台，输出 "hello world"
	t.Execute(os.Stdout, "world")
	*/
	//template.New("").Funcs(funcMap FuncMap):将所有的自定义函数添加到模板对象中。必须在解析模板之前调用它。
	//func (t *Template) ParseGlob(pattern string):根据给定的模板文件路径模式，解析符合模式的所有模板文件，并以模板名称为 key，模板内容为 value，将解析结果添加到当前的模板对象中。
	//template.Must:主要用于代码简化，避免在处理模板时频繁检查错误，简化代码逻辑。
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
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

	c.engine = engine //实例化时，需要给c.engine赋值

	engine.router.handle(c)
}
