package main

// 首先将两个数进行异或运算，得到一个新的二进制数，只需要统计该二进制数中1的个数即可
// 那么如何统计二进制数中1的个数呢？
// 答案是，和1进行与操作，如果一个数的末尾为1，则与的结果为1，计数器增加1，然后右移一位，依次进行与操作即可

func hammingDistance(x int, y int) int {
	newNum := x ^ y
	count := 0
	for newNum != 0 {
		if newNum & 1 == 1 {
			count++
		}
		newNum = newNum >> 1
	}
	return count
}

func main() {

}
