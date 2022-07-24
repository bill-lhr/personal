/**
 * @Author: lihaoran
 * @Description:
 * @Date: 2022/7/21 00:14
 */

package min_path_sum

import "github.com/bill_lhr/personal/common"

// https://leetcode.cn/problems/minimum-path-sum/
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[0][0]
			} else if i == 0 {
				dp[i][j] = dp[0][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][0] + grid[i][j]
			} else {
				dp[i][j] = common.Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}
