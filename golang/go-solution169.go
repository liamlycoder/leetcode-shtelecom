package main

/**
摩尔投票法：
1)假设当前数为众数, 遇到相同的数字则加一,否则减一
2)若前n个票和为0(互相抵消), 则设下一个数为当前数
3)重复1, 2 最后的当前数为众数
 */

func majorityElement(nums []int) int {
	vote := 0
	x := nums[0]
	for _, num := range nums {
		// 票数抵消
		if vote == 0 {
			x = num
		}
		if num == x {
			// 投正票
			vote++
		} else {
			// 投反票
			vote--
		}

	}
	return x
}

func main() {

}
