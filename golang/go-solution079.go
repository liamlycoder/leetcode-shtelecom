package main

type pair struct { x, y int }

var directions = []pair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	height, width := len(board), len(board[0])
	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}

	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}

		if k == len(word) - 1 {
			return true
		}

		visited[i][j] = true
		defer func() {visited[i][j] = false}()    // 回溯需要还原状态

		for _, dir := range directions {
			if nextI, nextJ := i + dir.x, j + dir.y; nextI >= 0 && nextI < height && nextJ >= 0 && nextJ < width && !visited[nextI][nextJ] {
				if check(nextI, nextJ, k + 1) {
					return true
				}
			}
		}

		return false
	}

	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}

	return false
}


func main() {

}
