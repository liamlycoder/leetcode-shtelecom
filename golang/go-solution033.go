package main

func search(nums []int, target int) int {
	if nums == nil {
		return -1
	}

	left, right := 0, len(nums) - 1

	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		}
		// 左边部分有序
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
			// 右边部分有序
		} else {
			if nums[mid] < target && target <= nums[len(nums) - 1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {

}
