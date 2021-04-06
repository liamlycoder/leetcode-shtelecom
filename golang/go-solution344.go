package main

func reverseString(s []byte)  {
	n := len(s)
	if n == 0 {
		return
	}
	left, right := 0, n - 1
	for left < right {
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		left++
		right--
	}
}

func main() {

}
