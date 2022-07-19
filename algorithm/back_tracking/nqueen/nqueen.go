/**
 * @Author: lihaoran04@baidu.com
 * @Date: 2020/9/13 10:21 下午
 */

package nqueen

func nQueens(n int) int {
	res, i := 0, 0 // i表示已放queen的行数
	board := make([][]bool, n)
	var backTracking func()
	backTracking = func() {
		// 递归退出条件，处理结果
		if i == n {
			res++
			return
		}

		// 遍历棋盘每一行
		board[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			// 检查当前位置是否能放
			if check(board, i, j) {
				board[i][j] = true
			} else {
				continue
			}
			i++
			backTracking()
			i--
			board[i][j] = false
		}
	}

	backTracking()
	return res
}

func check(board [][]bool, row, col int) bool {
	n := len(board)
	// 检查上方
	for i := row - 1; i >= 0; i-- {
		if board[i][col] {
			return false
		}
	}
	// 检查左上
	i, j := row-1, col-1
	for i >= 0 && j >= 0 {
		if board[i][j] {
			return false
		}
		i--
		j--
	}
	// 检查右上
	i, j = row-1, col+1
	for i >= 0 && j < n {
		if board[i][j] {
			return false
		}
		i--
		j++
	}
	return true
}
