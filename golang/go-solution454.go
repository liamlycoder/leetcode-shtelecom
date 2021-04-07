package main

func fourSumCount(A []int, B []int, C []int, D []int) int {
	record := make(map[int]int)
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			record[C[i] + D[j]]++
		}
	}

	res := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if _, ok := record[0 - A[i] - B[j]]; ok {
				res += record[0 - A[i] - B[j]]
			}
		}
	}
	return res
}

func main() {

}
