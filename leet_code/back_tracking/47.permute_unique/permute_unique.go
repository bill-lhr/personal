/**
 * @Author: lihaoran
 * @Description:
 * @Date: 2022/7/19 00:13
 */

package permute_unique

import "sort"

// https://leetcode.cn/problems/permutations-ii/
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	list := make([]int, 0)
	used := make(map[int]bool) // 使用过的元素下标
	var backTracking func()
	sort.Ints(nums)

	backTracking = func() {
		if len(list) == len(nums) {
			tmp := make([]int, len(list))
			copy(tmp, list)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			list = append(list, nums[i])
			backTracking()
			list = list[0 : len(list)-1]
			used[i] = false
		}
	}
	backTracking()
	return res
}
