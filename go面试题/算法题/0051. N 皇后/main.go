package _051__N_皇后

import "strings"

var (
	res [][]string
)

func solveNQueens(n int) [][]string {
	res = make([][]string, 0)
	chessboard := make([][]string, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
	}
	//初始化全为“.”
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}
	dfs(chessboard, n, 0)
	return res
}

//回溯函数
func dfs(chessboard [][]string, n, row int) {
	if row == n {
		tmp := make([]string, n)
		for i, val := range chessboard { //将【】string 转成string
			tmp[i] += strings.Join(val, "")
		}
		res = append(res, tmp)
	}

	for i := 0; i < n; i++ { //遍历一层
		if !isValid(chessboard, n, row, i) {
			continue
		}
		chessboard[row][i] = "Q"
		dfs(chessboard, n, row+1)
		chessboard[row][i] = "."
	}
}

//验证棋盘合法性函数
func isValid(chessboard [][]string, n, row, col int) bool {
	//判断同一列上方是否有皇后
	for i := 0; i < row; i++ {
		if chessboard[i][col] == "Q" {
			return false
		}
	}
	//判断向左45度，是否有皇后
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	//	判断向右45度，是否有皇后
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	return true
}
