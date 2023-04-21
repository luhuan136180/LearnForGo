package gee

import (
	"net/http"
	"strings"
)

/*      将 Trie 树应用到路由中
使用 roots 来存储每种请求方式的Trie 树根节点。使用 handlers 存储每种请求方式的 HandlerFunc 。
	getRoute 函数中，还解析了:和*两种匹配符的参数，返回一个 map 。
	例如/p/go/doc匹配到/p/:lang/doc，解析结果为：{lang: "go"}，/static/css/geektutu.css匹配到/static/*filepath，解析结果为{filepath: "css/geektutu.css"}
*/

type router struct {
	roots    map[string]*node       //存储每种请求方式的路由树的根节点
	handlers map[string]HandlerFunc //还是路由路径映射路由函数
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

//parse:解析   pattern:模式
func parsePattern(pattern string) []string {
	//用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item) //添加
			if item[0] == '*' {         //发现是 * ,退出，保证路径中只有一个 *
				break
			}
		}
	}
	return parts
}

//添加路由节点到路由树中
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok { //不存在该method的路由树，初始化
		r.roots[method] = &node{}
	}

	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler

}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path) //确定的路由路径，
	params := make(map[string]string)

	root, ok := r.roots[method] //root为查找的那个路由树的根节点
	if !ok {                    //该路由树没有注册，不可能有对应的路由方法，退出
		return nil, nil
	}

	n := root.search(searchParts, 0)
	if n != nil { //找到了匹配节点，其中包含有匹配路径等信息
		parts := parsePattern(n.pattern) //真正匹配到的路由路径，其中可能含有:或者*这两个动态匹配字符，可能与传参path不完全相同
		for index, part := range parts {
			if part[0] == ':' { //某一个路径段是动态路径，
				params[part[1:]] = searchParts[index] //传参，获取url路径中的实际值
			}
			//下面这段没分析懂
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		} //遍历完了，
		return n, params
	}

	//没有该匹配的节点，返回空
	return nil, nil
}

//没搞懂
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
	n, params := r.getRoute(c.Method, c.Path) //查询
	if n != nil {                             //将该次请求的内容填充进该次请求对应的自定义context中，方便取用
		c.Params = params //将解析出来的路由参数赋值给c.Params
		// 构造处理函数的 key，由请求方法和路由规则构成，用于从 handlers 字典中取出对应处理函数
		key := c.Method + "-" + n.pattern
		// 执行处理函数，将当前请求对应的上下文对象 c 作为参数传入
		r.handlers[key](c)
	} else {
		// 如果未找到对应的路由规则，则构造 404 响应，并返回到客户端
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
