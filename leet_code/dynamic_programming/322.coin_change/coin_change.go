package coin_change

import "math"

// https://leetcode.cn/problems/coin-change/
// eg: coins == [1,2,5] amount == 11

// 自顶向下 备忘录剪枝
// dp方法定义: dp(i) 表示金额为i，最少需要dp(i)个硬币
// 状态转义: dp(i) = min{ dp(i-coin) | coin ∈ coins} 注意--dp(i-coin)必须有解才能做min计算

func coinChange(coins []int, amount int) int {
	memo := make(map[int]int) // 备忘录
	var dp func(int) int

	dp = func(amount int) int {
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return -1
		}
		if _, ok := memo[amount]; ok {
			return memo[amount]
		}

		res := math.MaxInt32
		for _, coin := range coins {
			subRes := dp(amount - coin)
			if subRes == -1 {
				continue
			}
			res = min(res, subRes) + 1
		}
		if res == math.MaxInt32 {
			res = -1
		}
		memo[amount] = res
		return res
	}
	return dp(amount)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
