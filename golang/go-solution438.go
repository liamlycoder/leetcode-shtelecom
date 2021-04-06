package main

func findAnagrams(s string, p string) []int {
	window := make(map[byte]int)
	need := make(map[byte]int)
	left, right := 0, 0
	n := len(s)
	var res []int
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	for right < n {
		window[s[right]]++
		for window[s[right]] > need[s[right]] {
			window[s[left]]--
			left++
		}
		if right - left + 1 == len(p) {
			res = append(res, left)
		}
		right++
	}
	return res
}

func main() {

}
