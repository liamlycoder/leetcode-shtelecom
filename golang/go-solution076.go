package main

import "math"

func minWindow(s string, t string) string {
	ori, cnt := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	minLen := math.MaxInt64
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	for left, right := 0, 0; right < sLen; right++ {
		if right < sLen && ori[s[right]] > 0 {
			cnt[s[right]]++
		}
		for check() && left <= right {
			if (right - left + 1) < minLen {
				minLen = right - left + 1
				ansL, ansR = left, left + minLen
			}
			if _, ok := ori[s[left]]; ok {
				cnt[s[left]]--
			}
			left++
		}
	}

	if ansL == -1 {
		return ""
	}

	return s[ansL: ansR]
}

func main() {

}
