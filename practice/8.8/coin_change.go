package __8

import "github.com/bill_lhr/personal/common"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin && dp[i-coin] != -1 {
				if dp[i] == -1 {
					dp[i] = dp[i-coin] + 1
				} else {
					dp[i] = common.Min(dp[i], dp[i-coin]+1)
				}
			}
		}
	}
	return dp[amount]
}
