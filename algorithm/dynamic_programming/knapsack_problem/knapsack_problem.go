/**
 * @Author: lihaoran
 * @Description:
 * @Date: 2022/7/23 23:12
 */

package knapsack_problem

import "github.com/bill_lhr/personal/common"

/*
	1.dp数组定义: dp[i][j] 表示[0,i]中的物品放进容量为j的背包里能获得的最大价值
	2.递归方程: dp[i][j] = max{ dp[i-1][j], dp[i-1][j-weight[i-1]]+value[i] }
	3.初始化: dp[i][0]表示容量为0的背包，所以最大价值为0；dp[0][j]则表示容量为j只可放value[0]，若j>=weight[0]，则价值为value[0]
*/
func knapsack(weight, value []int, bagWeight int) int {
	n := len(weight)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, bagWeight+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= bagWeight; j++ {
			if j < weight[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = common.Max(dp[i-1][j], dp[i-1][j-weight[i-1]]+value[i-1])
			}
		}
	}

	return dp[n][bagWeight]
}
