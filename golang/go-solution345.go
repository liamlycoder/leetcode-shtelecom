package main


func reverseVowels(s string) string {
	s2 := []byte(s)
	n := len(s)
	if n == 0 {
		return s
	}
	left, right := 0, n - 1
	for left < right {
		for left < right && !isVowel(s2[left]) {
			left++
		}
		for left < right && !isVowel(s2[right]) {
			right--
		}
		if left < right {
			temp := s2[left]
			s2[left] = s2[right]
			s2[right] = temp
			left++
			right--
		}
	}
	return string(s2)
}

func isVowel(ch byte) bool {
	return ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' ||
		ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

func main() {

}
