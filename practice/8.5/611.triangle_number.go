package __5

import "sort"

func triangleNumber(nums []int) int {
	cnt := 0
	if len(nums) < 3 {
		return cnt
	}
	sort.Ints(nums)
	// 固定第一条边
	for i := len(nums) - 1; i > 1; i-- {
		l, r := 0, i-1
		for l < r {
			if nums[l]+nums[r] > nums[i] {
				cnt += r - l
				r--
			} else {
				l++
			}
		}
	}
	return cnt
}
