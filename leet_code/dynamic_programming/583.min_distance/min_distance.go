package min_distance

import "github.com/bill-lhr/personal/common"

// https://leetcode.cn/problems/delete-operation-for-two-strings/

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = common.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	lcs := dp[m][n]
	return m - lcs + n - lcs
}
