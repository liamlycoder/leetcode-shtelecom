package main

import "fmt"

func isValid(s string) bool {
	n := len(s)
	if n % 2 == 1 {
		return false
	}
	pairs := map[byte]byte {
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []byte
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {   // 大于0说明出现了右括号
			if len(stack) == 0 || stack[len(stack) - 1] != pairs[s[i]] {   // 栈顶元素不匹配，直接返回false
				return false
			}
			stack = stack[:len(stack) - 1]  // 如果能匹配，将栈顶元素出栈
		} else {  // 否则的话，说明出现的是左括号，只需要加入栈即可
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}

func main() {
	pairs := map[byte]byte {
		')': '(',
		']': '[',
		'}': '{',
	}
	fmt.Println(pairs[')'])

}
