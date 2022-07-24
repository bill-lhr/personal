/**
 * @Author: lihaoran
 * @Date: 2022/7/24 13:08
 */

package three_sum

import (
	"sort"
)

/**
 * @Description:
	给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
	注意：答案中不可以包含重复的三元组。
	链接：https://leetcode.cn/problems/3sum
*/
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	// 排序
	sort.Sort(sort.IntSlice(nums))
	res := make([][]int, 0)
	// 确定第一个数
	for i := 0; i < len(nums); i++ {
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		first := nums[i]
		// 当first > 0时，因为已经排序，因此之后两数都大于0，无法得到和为0
		if first > 0 {
			break
		}

		low, high := i+1, len(nums)-1
		for low < high {
			if first+nums[low]+nums[high] == 0 {
				res = append(res, []int{first, nums[low], nums[high]})
				// 去重
				for low < high && nums[low] == nums[low+1] {
					low++
				}
				for low < high && nums[high] == nums[high-1] {
					high--
				}
				low++
				high--
			} else if first+nums[low]+nums[high] < 0 {
				low++
			} else {
				high--
			}
		}
	}

	return res
}
