package _150__逆波兰表达式求值

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for i := 0; i < len(tokens); i++ {
		val, err := strconv.Atoi(tokens[i])
		if err == nil {
			stack = append(stack, val)
		} else {
			//遇到符号
			//先弹出后两个数字
			num2, num1 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch tokens[i] {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}

	}

	return stack[0]

}
