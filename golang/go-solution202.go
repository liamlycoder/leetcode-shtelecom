package main

// 哈希表的方法
func isHappy(n int) bool {
	m := map[int]bool{}
	for ; n != 1 && !m[n]; n, m[n] = step(n), true {}
	return n == 1
}

// 快慢指针的方法，相比哈希表而言，空间复杂度从O(n)降低为O(1)  时间复杂度都为O(logn)
func isHappy2(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

func main() {

}
