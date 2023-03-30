package _037__解数独

func solveSudoku(board [][]byte) {
	var backtracking func(board [][]byte) bool
	backtracking = func(board [][]byte) bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				//判断此位置是否适合填数字
				if board[i][j] != '.' {
					continue
				}
				//尝试填1-9
				for k := '1'; k <= '9'; k++ {
					if IsValid(i, j, byte(k), board) == true { //如果满足要求就填
						board[i][j] = byte(k)
						if backtracking(board) == true {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
		return true
	}
	backtracking(board)
}

//判断填入数字是否满足要求
func IsValid(row, col int, k byte, board [][]byte) bool {
	//	判断一行没有重复匀速
	for i := 0; i < 9; i++ {
		if board[row][i] == k {
			return false
		}
	}
	//	判断一列没有重复元素
	for i := 0; i < 9; i++ {
		if board[i][col] == k {
			return false
		}
	}
	//	判断小九宫格没有重复元素
	startrow := (row / 3) * 3
	startcol := (col / 3) * 3
	for i := startrow; i < startrow+3; i++ {
		for j := startcol; j < startcol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}
