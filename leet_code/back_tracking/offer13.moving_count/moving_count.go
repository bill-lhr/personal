/**
 * @Author: lihaoran
 * @Description:
 * @Date: 2022/7/21 22:44
 */

package moving_count

// https://leetcode.cn/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
func movingCount(m int, n int, k int) int {
	board := make([][]bool, m)
	for i := range board {
		board[i] = make([]bool, n)
	}
	cnt := 0
	var dfs func(int, int)

	dfs = func(row, col int) {
		if row < 0 || row >= m || col >= n || col < 0 {
			return
		}
		// 可进入 && 没有进入过
		if check(row, col, k) && !board[row][col] {
			board[row][col] = true
			cnt++
		} else {
			return
		}
		// 上
		dfs(row-1, col)
		// 下
		dfs(row+1, col)
		// 左
		dfs(row, col-1)
		// 右
		dfs(row, col+1)
	}

	dfs(0, 0)
	return cnt
}

func check(i, j, k int) bool {
	sum := 0
	for i > 0 {
		sum += i % 10
		i /= 10
	}
	for j > 0 {
		sum += j % 10
		j /= 10
	}
	return sum <= k
}
