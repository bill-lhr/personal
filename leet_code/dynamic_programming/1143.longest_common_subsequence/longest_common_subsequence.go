package longest_common_subsequence

import (
	"github.com/bill_lhr/personal/common"
)

// https://leetcode.cn/problems/longest-common-subsequence/
/*
	输入：text1 = "abcde", text2 = "ace"
	输出：3
	解释：最长公共子序列是 "ace" ，它的长度为 3 。

	dp定义：dp(i,j) 表示text1[0:i]、text2[0:j]的lcs
	状态转移：
		if text1[i] == text2[j]
			dp(i,j) = dp(i-1,j-1) + 1
		else
			dp(i,j) = max{ dp(i-1,j), dp(i,j-1) }
	结果：dp(len(text1)-1,len(text2)-1)
*/

// 自顶向下
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	memo := make([][]int, m)
	for i := range memo {
		for j := 0; j < n; j++ {
			memo[i] = append(memo[i], -1)
		}
	}

	var dp func(int, int) int
	dp = func(i int, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		if text1[i] == text2[j] {
			memo[i][j] = dp(i-1, j-1) + 1
			return memo[i][j]
		} else {
			memo[i][j] = common.Max(dp(i-1, j), dp(i, j-1))
			return memo[i][j]
		}
	}

	return dp(m-1, n-1)
}

// 自底向上
func longestCommonSubsequence2(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	// 初始化dp数组
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = common.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
