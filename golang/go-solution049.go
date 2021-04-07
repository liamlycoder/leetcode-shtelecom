package main

func groupAnagrams(strs []string) [][]string {
	// 数组为键，对应的字符串数组为值
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b - 'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}

	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}

	return ans
}

func main() {

}
