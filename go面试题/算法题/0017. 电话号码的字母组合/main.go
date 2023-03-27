package _017__电话号码的字母组合

//字母数量为深度——也就是递归层数
//解决数字和字母如何映射

//需要一个总结果集，一个单一结果集
var ( //需要在函数中初始化
	res   []string
	path  []byte
	m     []string
	depth int
)

func letterCombinations(digits string) []string {
	m = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res, path = make([]string, 0), make([]byte, 0)
	depth = len(digits)
	index := 0
	if digits == "" {
		return res
	}
	dfs(digits, index)
	return res
}

func dfs(digits string, index int) {
	if len(path) == depth {
		res = append(res, string(path)) //加入结果集
		return
	}
	tmp := int(digits[index] - '0') //获取digits[index]位置的数字
	str := m[tmp-2]                 //获取该数字对应的字母,从数字2开始，要注意对其
	for _, val := range str {       //遍历str
		path = append(path, byte(val))
		dfs(digits, index+1)      //递归
		path = path[:len(path)-1] //回溯
	}
}
