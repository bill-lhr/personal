package __11

import "sort"

func triangleNumber(nums []int) int {
	res := 0
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		left, right := 0, i-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum > nums[i] {
				res += right - left
				right--
			} else {
				left++
			}
		}
	}
	return res
}
