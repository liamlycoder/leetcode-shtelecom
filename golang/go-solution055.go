package main

func canJump(nums []int) bool {
	res := 0
	for k, v := range nums {
		if k > res {
			return false
		}
		res = max(res, k + v)
	}
	return true
}

func main() {

}
