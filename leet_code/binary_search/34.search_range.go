package binary_search

/**
 * @Description:
	给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
	如果数组中不存在目标值 target，返回[-1, -1]。
	你必须设计并实现时间复杂度为O(log n)的算法解决此问题。
	链接：https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array
*/
func searchRange(nums []int, target int) []int {
	low, high, targetIdx := 0, len(nums)-1, -1
	// 二分查找重点！！！ 小于等于
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			targetIdx = mid
			break
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if targetIdx == -1 {
		return []int{-1, -1}
	}
	low, high = targetIdx, targetIdx
	for low >= 0 && nums[low] == target {
		low--
	}
	for high < len(nums) && nums[high] == target {
		high++
	}
	return []int{low + 1, high - 1}
}
