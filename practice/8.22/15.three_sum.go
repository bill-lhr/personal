package __22

import "sort"

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if i < len(nums)-1 && nums[i] == nums[i+1] {
			continue
		}
		left, right := 0, i-1
		for left < right {
			if left > 0 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < i-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			sum := nums[left] + nums[right] + nums[i]
			if sum == 0 {
				res = append(res, []int{nums[left], nums[right], nums[i]})
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
