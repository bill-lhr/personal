/**
 * @Author: lihaoran
 * @Date: 2022/7/25 23:40
 * @Description:
 */

package _725

import (
	"github.com/bill_lhr/personal/common"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = -1
	}

	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin == 0 {
				dp[i] = 1
				break
			} else if i-coin > 0 && dp[i-coin] != -1 {
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
