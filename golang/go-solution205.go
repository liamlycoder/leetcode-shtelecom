package main

func isIsomorphic(s string, t string) bool {
	hash1, hash2 := make(map[byte]byte), make(map[byte]byte)
	if len(s) != len(t) {
		return false
	}

	for i := range s {
		x, y := s[i], t[i]
		if hash1[x] > 0 && hash1[x] != y || hash2[y] > 0 && hash2[y] != x {
			return false
		}
		hash1[x] = y
		hash2[y] = x
	}

	return true
}

func main() {

}
