/**
 * @Author: lihaoran04@baidu.com
 * @Date: 2020/9/13 10:21 下午
 */

package main

import "fmt"

const N = 8

var board [8][8]bool
var count int

func check(x, y int) bool {
	// 行列检查
	for i := 0; i < N; i++ {
		if board[i][y] || board[x][i] {
			return false
		}
	}
	// 对角线检查
	for i := 1; x-i >= 0 && y-i >= 0; i++ {
		if board[x-i][y-i] {
			return false
		}
	}
	for i := 1; x-i >= 0 && y+i < N; i++ {
		if board[x-i][y+i] {
			return false
		}
	}
	return true
}

func dfs(x, y int) {
	if x == N {
		count++
		return
	}
	for i := 0; i < N; i++ {
		if check(x, i) && !board[x][i] {
			board[x][i] = true
			dfs(x+1, i)
			board[x][i] = false
		}
	}
}

func main() {
	dfs(0, 0)
	fmt.Println(count)
}
