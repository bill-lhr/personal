/**
 * @Author: lihaoran
 * @Date: 2022/7/24 14:29
 */

package four_sum

import "sort"

/*
* @Description:
	给你一个由 n 个整数组成的数组nums ，和一个目标值 target 。
	请你找出并返回满足下述全部条件且不重复的四元组[nums[a], nums[b], nums[c], nums[d]]（若两个四元组元素一一对应，则认为两个四元组重复）：
		0 <= a, b, c, d< n
		a、b、c 和 d 互不相同
		nums[a] + nums[b] + nums[c] + nums[d] == target
	你可以按 任意顺序 返回答案 。
	链接：https://leetcode.cn/problems/4sum
*/
func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return nil
	}
	sort.Sort(sort.IntSlice(nums))

	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		first := nums[i]
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			second := nums[j]
			if j-1 > i && nums[j] == nums[j-1] {
				continue
			}
			low, high := j+1, n-1
			for low < high {
				sum := first + second + nums[low] + nums[high]
				if sum < target {
					low++
				} else if sum > target {
					high--
				} else {
					res = append(res, []int{first, second, nums[low], nums[high]})
					// 去重
					for low < high && nums[low] == nums[low+1] {
						low++
					}
					for low < high && nums[high] == nums[low-1] {
						high--
					}
					low++
					high--
				}
			}
		}
	}

	return res
}
