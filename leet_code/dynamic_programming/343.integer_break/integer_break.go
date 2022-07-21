/**
 * @Author: lihaoran
 * @Description:
 * @Date: 2022/7/21 23:09
 */

package integer_break

import "github.com/bill_lhr/personal/common"

// https://leetcode.cn/problems/integer-break/
// 给定一个正整数 n，将其拆分为 k 个正整数的和（ k >= 2 ），并使这些整数的乘积最大化。
/*
	dp定义：dp[i]表示正整数i的拆分最大乘积
	分析：1.假设拆分第一个正整数j（1 <= j < i），则有：1) k = 2，有可能的解 j * (i-j)；2) k > 2，有可能解 j * dp[i-j]
		 2.因为 dp[i-j] 为 i-j 拆分后的最大乘积，因此当拆分第一个数为j时，最优解为 max{ j * (i-j), j * dp[i-j] }
	dp方程: dp[i] = max{ max{ j * (i-j), j * dp[i-j] } | 1 <= j < i }
*/
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1

	for i := 2; i <= n; i++ {
		max := 0
		for j := 1; j < i; j++ {
			tmp := common.Max(j*(i-j), j*dp[i-j])
			max = common.Max(max, tmp)
		}
		dp[i] = max
	}

	return dp[n]
}
