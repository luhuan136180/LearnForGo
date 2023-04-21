package gee

import (
	"fmt"
	"strings"
)

/*
接下来我们实现的动态路由具备以下两个功能。

参数匹配:。例如 /p/:lang/doc，可以匹配 /p/c/doc 和 /p/go/doc。
通配*。例如 /static/*filepath，可以匹配/static/fav.ico，也可以匹配/static/js/jQuery.js，这种模式常用于静态服务器，能够递归地匹配子路径。
*/
type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang,是全路径？
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

//为了实现动态路由匹配，加上了isWild这个参数。
//即当我们匹配 /p/go/doc/这个路由时，第一层节点，p精准匹配到了p，第二层节点，go模糊匹配到:lang，那么将会把lang这个参数赋值为go，继续下一层匹配。

//matchChild函数用于在node的子节点中寻找与part匹配的节点，如果不存在则返回nil。
//返回第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node { //用于插入
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

//matchChildren函数则用于在node的子节点中寻找所有与part匹配的节点，返回一个节点数组。
//返回匹配成功的所有节点，用于查找？
func (n *node) matchChildren(part string) []*node { //只用于单层
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

/*
对于路由来说，最重要的当然是注册与匹配了。开发服务时，注册路由规则，映射handler；访问时，匹配路由规则，查找到对应的handler。因此，Trie 树需要支持节点的插入与查询。

插入功能很简单，递归查找每一层的节点，如果没有匹配到当前part的节点，则新建一个，有一点需要注意，
/p/:lang/doc只有在第三层节点，即doc节点，pattern才会设置为/p/:lang/doc。p和:lang节点的pattern属性皆为空。
因此，当匹配结束时，我们可以使用n.pattern == ""来判断路由规则是否匹配成功。例如，/p/python虽能成功匹配到:lang，但:lang的pattern值为空，因此匹配失败。
*/

//插入
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height { //路径遍历完毕，为叶子节点
		n.pattern = pattern //将匹配的路由全路径赋值给节点的参数
		return              //exit
	}

	part := parts[height]       //匹配当前的一小段，以 / 分割
	child := n.matchChild(part) //查询当前节点的子节点是否已经注册过相同字路径段
	if child == nil {           //没有注册过
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'} //初始化子节点
		n.children = append(n.children, child)                              //将初始化的子节点注册到当前节点的孩子节点列表里
	}

	//对孩子节点继续搜索遍历
	child.insert(pattern, parts, height+1) //递归
}

//查找
func (n *node) search(parts []string, height int) *node {
	// 如果路由路径已经匹配完成，或者当前节点的part为 * 后缀，则判断当前节点是否有pattern（路由规则），有则返回该节点，否则说明该路径无法匹配，返回nil
	//func HasPrefix(s, prefix string) bool:判断s是否有前缀字符串prefix。
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]             //获取当前的路径段
	children := n.matchChildren(part) //查找当前node的孩子节点中是否用路径段相同的子节点，返回切片

	// 递归调用所有子节点的search函数，并且传入下一级的路由路径部分和高度（递归深度），如果返回的节点不为空，就直接返回此节点
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	// 如果所有子节点都没有找到匹配的节点，则返回nil
	return nil
}

//没搞懂
func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}
