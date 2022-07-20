/**
 * @Author: lihaoran
 * @Description:
 * @Date: 2022/7/20 23:43
 */

package length_of_LIS

// https://leetcode.cn/problems/longest-increasing-subsequence/
func lengthOfLIS(nums []int) int {
	dp := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		dp = append(dp, 1)
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
