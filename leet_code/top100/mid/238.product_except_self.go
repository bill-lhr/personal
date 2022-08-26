package mid

/**
给你一个整数数组nums，返回 数组answer，其中answer[i]等于nums中除nums[i]之外其余各元素的乘积。
题目数据 保证 数组nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内。
请不要使用除法，且在O(n) 时间复杂度内完成此题。

链接：https://leetcode.cn/problems/product-of-array-except-self
*/
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			res[i] = 1
			continue
		}
		res[i] = res[i-1] * nums[i-1]
	}
	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			continue
		}
		right *= nums[i+1]
		res[i] = res[i] * right
	}
	return res
}
