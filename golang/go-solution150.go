package main

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {  // 说明字符串转数字成功，即碰到的是数字，那么我们直接入栈
			stack = append(stack, val)
		} else {  // 否则，说明遇到的是符号，那么我们弹出栈顶的两个元素，按照该符号进行运算，再把运算结果入栈
			num1, num2 := stack[len(stack) - 2], stack[len(stack) - 1]
			stack = stack[:len(stack) - 2]
			switch token {
			case "+":
				stack = append(stack, num1 + num2)
			case "-":
				stack = append(stack, num1 - num2)
			case "*":
				stack = append(stack, num1 * num2)
			case "/":
				stack = append(stack, num1 / num2)
			}
		}
	}
	return stack[0]
}
