package main

func removeDuplicates(nums []int) int {
	i, j := 0, 1
	n := len(nums)
	for i < n && j < n {
		if nums[i] == nums[j] {
			j++
		} else {
			i++
			nums[i] = nums[j]
			j++
		}
	}
	return i + 1
}

func main() {

}
