package __11

import "github.com/bill_lhr/personal/common"

func maximalSquare(matrix [][]byte) int {
	maxLen := 0
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1) // dp表示以dp[i][j]为右下角的最大正方形边长
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = common.Min(dp[i-1][j], common.Min(dp[i-1][j-1], dp[i][j-1])) + 1
			maxLen = common.Max(maxLen, dp[i][j])
		}
	}
	return maxLen * maxLen
}
