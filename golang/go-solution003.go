package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	hashTable := map[byte]int{}
	result, left := 0, 0
	for index, ch := range s {
		if repeat, ok := hashTable[byte(ch)]; ok {
			left = max(left, repeat + 1)
		}
		hashTable[byte(ch)] = index
		result = max(result, index - left + 1)
	}
	return result
}


func main() {
	s := "abcdefg"
	for i, j := range s {
		fmt.Println("i = ", i, "j = ", string(j))
	}

}
