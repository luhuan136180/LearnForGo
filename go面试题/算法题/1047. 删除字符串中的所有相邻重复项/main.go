package _047__删除字符串中的所有相邻重复项

func removeDuplicates(s string) string {
	var stack []byte

	for i := 0; i < len(s); i++ {
		//stack[len(stack)-1]表示最后一个元素
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else { // 栈不空 且 与栈顶元素不等 入栈
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
