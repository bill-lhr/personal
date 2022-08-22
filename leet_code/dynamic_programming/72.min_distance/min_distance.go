package min_distance

import "github.com/bill_lhr/personal/common"

/**
  * @Description:
	给你两个单词word1 和word2， 请返回将word1转换成word2 所使用的最少操作数 。
	你可以对一个单词进行如下三种操作：
	插入一个字符
	删除一个字符
	替换一个字符
	链接：https://leetcode.cn/problems/edit-distance

	1. dp[i][j] 表示word1[0:i] 和 word2[0:j]的最小编辑距离
*/

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 字符相同的时候就等于dp[i-1][j-1]
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 不同时
				// dp[i-1][j-1] 表示把word1[i]替换; dp[i-1][j] 表示word1[i]删除; dp[i][j-1] 表示在word[i]后边插入
				dp[i][j] = common.Min(dp[i-1][j-1], common.Min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}
	return dp[m][n]
}
