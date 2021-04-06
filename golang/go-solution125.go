package main

import "strings"

func isPalindrome2(s string) bool {
	s = strings.ToLower(s)
	n := len(s)
	if n == 0 {
		return true
	}

	left, right := 0, n - 1
	for left < right {
		for left < right && !isalnum(s[left]) {
			left++
		}
		for left < right && !isalnum(s[right]) {
			right--
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch > 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func main() {

}
