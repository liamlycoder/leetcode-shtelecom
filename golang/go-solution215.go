package main

func findKthLargest(nums []int, k int) int {
	left, right := 0, len(nums) - 1
	for {
		if left >= right {
			return nums[right]
		}
		p := partition(nums, left, right)
		if p + 1 == k {
			return nums[p]
		} else if p + 1 < k {
			left = p + 1
		} else {
			right = p - 1
		}
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	for cur := left; cur < right; cur++ {
		if nums[cur] > pivot {
			nums[cur], nums[left] = nums[left], nums[cur]
			left++
		}
	}
	nums[right], nums[left] = nums[left], nums[right]
	return left
}


func main() {

}
