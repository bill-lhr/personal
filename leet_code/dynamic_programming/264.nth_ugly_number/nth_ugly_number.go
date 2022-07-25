package nth_ugly_number

import (
	"github.com/bill-lhr/personal/common"
	"math"
)

/**
 * @Description:
	给你一个整数 n ，请你找出并返回第 n 个 丑数 。
	丑数 就是只包含质因数 2、3 和/或 5 的正整数。
	https://leetcode.cn/problems/ugly-number-ii/
*/
// dp[i]为 dp[1]~dp[i-1]分别乘{2,3,5}所得结果中最小的
func nthUglyNumber1(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt32
		for j := 1; j < i; j++ {
			for _, num := range []int{2, 3, 5} {
				tmp := dp[j] * num
				if tmp > dp[i-1] {
					dp[i] = common.Min(dp[i], tmp)
				}
			}
		}
	}
	return dp[n]
}

// 方法1存在大量重复计算，时间复杂度高，考虑只循环一次
// 用p2表示通过乘2获得了丑数的已有丑数下标，每个丑数都有一次分别乘2、3、5获得丑数的机会
// 因此可以用三个向右移动的下标来获得所有丑数
func nthUglyNumber2(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1

	for i := 2; i <= n; i++ {
		n2, n3, n5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = common.Min(n2, common.Min(n3, n5))
		if dp[i] == n2 {
			p2++
		}
		if dp[i] == n3 {
			p3++
		}
		if dp[i] == n5 {
			p5++
		}
	}
	return dp[n]
}
