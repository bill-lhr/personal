package coin_change

import "math"

// https://leetcode.cn/problems/coin-change/
// eg: coins == [1,2,5] amount == 11

// 1. 自顶向下 备忘录剪枝
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

		res := math.MaxInt
		for _, coin := range coins {
			subRes := dp(amount - coin)
			if subRes == -1 {
				continue
			}
			res = min(res, subRes) + 1
		}
		if res == math.MaxInt {
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
