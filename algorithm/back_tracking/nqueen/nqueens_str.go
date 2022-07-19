package nqueen

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	board := make([][]byte, 0, n)
	row := 0
	var backTracking func()

	// 初始化board
	for i := 0; i < n; i++ {
		tmp := make([]byte, 0, n)
		for j := 0; j < n; j++ {
			tmp = append(tmp, '.')
		}
		board = append(board, tmp)
	}

	backTracking = func() {
		if row == n {
			// 递归结束 记录结果
			s := make([]string, 0)
			for _, list := range board {
				s = append(s, string(list))
			}
			res = append(res, s)
			return
		}

		// 遍历每层
		for col := 0; col < n; col++ {
			if checkcheck(board, row, col) {
				board[row][col] = 'Q'
			} else {
				continue
			}
			row++
			backTracking()
			// 回溯还原
			row--
			board[row][col] = '.'
		}
	}

	backTracking()
	return res
}

func checkcheck(board [][]byte, row, col int) bool {
	n := len(board)
	// 检查上边
	for i := row - 1; i >= 0; i-- {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 检查左上
	i, j := row-1, col-1
	for i >= 0 && j >= 0 {
		if board[i][j] == 'Q' {
			return false
		}
		i, j = i-1, j-1
	}
	// 检查右上
	i, j = row-1, col+1
	for i >= 0 && j < n {
		if board[i][j] == 'Q' {
			return false
		}
		i, j = i-1, j+1
	}

	return true
}
