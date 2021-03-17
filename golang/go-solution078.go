package main

func subsets(nums []int) [][]int {
	n := len(nums)
	var result [][]int
	for mask := 0; mask < 1<<n; mask++ {
		var set []int
		for i, v := range nums {
			if mask>>i & 1 > 0 {
				set = append(set, v)
			}
		}
		result = append(result, set)
	}
	return result
}

func main() {

}
