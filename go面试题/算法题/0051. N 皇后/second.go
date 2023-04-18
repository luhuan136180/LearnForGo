package _051__N_皇后

import "strings"

var (
	//res        [][]string
	chessboard [][]string
)

func solveNQueens2(n int) [][]string {
	res, chessboard = make([][]string, 0), make([][]string, n)
	for i, _ := range chessboard {
		chessboard[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}
	dfs2(n, 0)

	return res
}

func dfs2(n, row int) {
	if row == n {
		//将一个棋盘（二维）转为一个一维string
		str := make([]string, n)
		for i, val := range chessboard {
			str[i] += strings.Join(val, "") //这里用断点理解一下，目前不是很清晰
		}
		res = append(res, str)
	}

	for col := 0; col < n; col++ {
		if istrue(chessboard, n, row+1, col) {
			chessboard[row][col] = "Q"
			dfs2(n, row+1)
			chessboard[row][col] = "."
		}
	}
}

func istrue(chessboard [][]string, n, row, col int) bool {
	// 同行没有q
	for i := 0; i < col; i++ {
		if chessboard[row-1][i] == "Q" {
			return false
		}
	}
	// 同列没有q
	for j := 0; j < row-1; j++ {
		if chessboard[j][col] == "Q" {
			return false
		}
	}
	// 左45度没有q
	for i, j := row-2, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	// 右45度没有q
	for i, j := row-2, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}

	return true
}
