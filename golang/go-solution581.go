package main

func findUnsortedSubarray(nums []int) int {
	mxv := nums[0]
	miv := nums[len(nums) - 1]
	j := len(nums) - 1
	l, r := 0, 0
	for i, v := range nums {
		if v < mxv {
			l = i
		} else {
			mxv = v
		}

		if v > miv {
			r = j
		} else {
			miv = nums[j]
		}
		j--
	}

	if l >= r {
		return 0
	}
	return r - l + 1
}
