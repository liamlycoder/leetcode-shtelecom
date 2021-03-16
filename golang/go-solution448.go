package main

func findDisappearedNumbers(nums []int) []int {
	var result []int
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n   // 当我们遍历到某个位置时，其中的数可能已经被增加过，因此需要对n取模来还原出它本来的值。
		nums[v] += n
	}

	for i, v := range nums {
		if v <= n {
			result = append(result, i + 1)
		}
	}
	return result
}

func main() {

}
