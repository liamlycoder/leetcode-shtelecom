package main


func maxArea(height []int) int {
	l, r := 0, len(height) - 1
	result := 0
	for l < r {
		area := min(height[l], height[r]) * (r - l)
		result = max(area, result)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return result
}

func main() {

}
