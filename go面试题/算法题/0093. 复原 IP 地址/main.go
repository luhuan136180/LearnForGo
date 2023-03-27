package _093__复原_IP_地址

import (
	"strconv"
	"strings"
)

var (
	res []string
	ip  []string
)

func restoreIpAddresses(s string) []string {
	startindex := 0
	//一定要记得初始化全局变量
	res, ip = make([]string, 0), make([]string, 0)
	dfs(s, startindex)
	return res
}

func dfs(s string, startindex int) {
	if len(ip) == 4 { //已经切割成四份
		if startindex == len(s) { //确定已经遍历完成，没有还没分配的数字
			str := strings.Join(ip, ".")
			res = append(res, str)
		}
		return
	}

	for i := startindex; i < len(s); i++ {
		if i != startindex && s[startindex] == '0' { //当前str不止一位数，其起始不为0
			break
		}
		str := s[startindex : i+1]
		num, _ := strconv.Atoi(str)
		if num >= 0 && num <= 255 {
			ip = append(ip, str)
			dfs(s, i+1)
			ip = ip[:len(ip)-1]
		} else {
			break
		}
	}

}
