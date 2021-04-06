package main

func removeDuplicates2(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	i, j := 2, 2
	for i < n && j < n {
		if nums[i - 2] != nums[j] {
			nums[i] = nums[j]
			i++
			j++
		} else {
			j++
		}
	}
	return i
}

func main() {

}
