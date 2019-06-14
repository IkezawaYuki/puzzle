package main

import "fmt"

func noConflict(board [][]int, current int, qindex int, n int) bool {
	for j := 0; j < current; j++ {
		if board[qindex][j] == 1 {
			return false
		}
	}
	var k int = 1
	for {
		if qindex-k >= 0 || current-k >= 0 {
			break
		}
		if board[qindex-k][current-k] == 1 {
			return false
		}
		k++
	}
	k = 1
	for {
		if qindex+k >= 0 && current-k >= 0 {
			break
		}
		if board[qindex+k][current-k] == 1 {
			return false
		}
		k++
	}
	return true
}

func main() {
	board := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	fmt.Println(len(board))
	n := len(board)
	for _, rows := range board {
		fmt.Println(rows)
		for _, cells := range rows {
			fmt.Println(cells)
		}
	}

	for i := 0; i < n; i++ {
		board[i][0] = 1
		for j := 0; j < n; j++ {
			board[j][1] = 1
			if noConflict(board, 1, j, n) {
				for k := 0; k < n; k++ {
					board[k][2] = 1
					if noConflict(board, 2, k, n) {
						for l := 0; l < n; l++ {
							board[l][3] = 1
							if noConflict(board, 3, l, n) {
								for _, row := range board {
									fmt.Println(row)
								}
								fmt.Println("-----")
							}
							board[l][3] = 0
						}
					}
					board[k][2] = 0
				}
			}
			board[j][1] = 0
		}
		board[i][0] = 0
	}
}
