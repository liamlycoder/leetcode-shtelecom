package main

func nextPermutation(nums []int)  {
	// 长度小于1就不用管了
	if len(nums) <= 1 {
		return
	}

	i, j, k := len(nums) - 2, len(nums) - 1, len(nums) - 1

	// 从后向前查找，找到第一个nums[i] < nums[j]的即可停止
	// 这里尤其要注意，当i等于0的时候也是需要判断的
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 {
		// 从后往前找出第一个满足nums[k] > nums[i]的k，然后交换nums[i]和nums[k]
		for k >= 0 && nums[k] <= nums[i] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}

	// 将j到结尾的部分逆置一下
	for x, y := j, len(nums) - 1; x < y; x, y = x + 1, y - 1 {
		nums[x], nums[y] = nums[y], nums[x]
	}
}

func main() {

}
