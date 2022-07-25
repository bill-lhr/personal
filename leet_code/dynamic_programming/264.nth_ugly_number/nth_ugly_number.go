package nth_ugly_number

import "github.com/bill-lhr/personal/common"

/**
 * @Description:
	给你一个整数 n ，请你找出并返回第 n 个 丑数 。
	丑数 就是只包含质因数 2、3 和/或 5 的正整数。
	https://leetcode.cn/problems/ugly-number-ii/
*/
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			for _, num := range []int{2, 3, 5} {
				dp[i] = common.Min(dp[i], dp[j]*num)
			}
		}
	}
	return dp[n]
}
