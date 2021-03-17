package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	p1, p2 := nums[0], max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		p1, p2 = p2, max(p1 + nums[i], p2)
	}

	return p2
}

func main() {

}
