/**
 * @Author: lihaoran
 * @Date: 2022/8/21 19:49
 */

package hard

import "github.com/bill_lhr/personal/common"

/**
 * @Description: 72.编辑距离
	给你两个单词word1 和word2， 请返回将word1转换成word2 所使用的最少操作数。
	你可以对一个单词进行如下三种操作：
	插入一个字符
	删除一个字符
	替换一个字符
	链接：https://leetcode.cn/problems/edit-distance
*/

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = common.Min(dp[i-1][j-1], common.Min(dp[i-1][j], dp[i][j-1]))
			}
		}
	}

	return dp[m][n]
}
