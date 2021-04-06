package main

func removeElement(nums []int, val int) int {
	i, j := 0, 0
	n := len(nums)
	for i < n && j < n {
		if nums[j] == val {
			j++
		} else {
			nums[i] = nums[j]
			i++
			j++
		}
	}
	return i
}

func main() {

}
