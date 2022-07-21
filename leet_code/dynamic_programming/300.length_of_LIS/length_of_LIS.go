/**
 * @Author: lihaoran
 * @Description:
 * @Date: 2022/7/20 23:43
 */

package length_of_LIS

import (
	"github.com/bill_lhr/personal/common"
)

// https://leetcode.cn/problems/longest-increasing-subsequence/

// 自底向上
func lengthOfLIS(nums []int) int {
	dp := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		dp = append(dp, 1)
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = common.Max(dp[i], dp[j]+1)
			}
		}
	}

	res := 0
	for _, n := range dp {
		if n > res {
			res = n
		}
	}
	return res
}

// 自顶向下
func lengthOfLIS2(nums []int) int {
	// dp函数 dp(i)表示以i结尾的最长递增子序列长度
	var dp func(int) int
	// 备忘录默认值为-1 标记该点没有处理过 结果为备忘录中最大值
	memo := make([]int, len(nums))
	for i := range nums {
		memo[i] = -1
	}

	dp = func(i int) int {
		if memo[i] != -1 {
			return memo[i]
		}
		tmpMax := 0
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				tmpMax = common.Max(tmpMax, dp(j))
			}
		}
		return tmpMax + 1
	}

	res := 0
	for i := range nums {
		memo[i] = dp(i)
		res = common.Max(res, memo[i])
	}

	return res
}
