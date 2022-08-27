/**
 * @Author: lihaoran
 * @Date: 2022/8/27 15:40
 */

package binary_search

/**
 * @Description: 已知存在一个按非降序排列的整数数组 nums ，数组中的值不必互不相同。
	给你 旋转后 的数组 nums 和一个整数 target ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 nums 中存在这个目标值 target ，则返回 true ，否则返回 false 。
	链接：https://leetcode.cn/problems/search-in-rotated-sorted-array-ii
*/
func search2(nums []int, target int) bool {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[j] {
			j--
		} else if nums[mid] < nums[j] {
			if target > nums[mid] && target <= nums[j] {
				i = mid + 1
			} else {
				j = mid - 1
			}
		} else {
			if target >= nums[i] && target < nums[mid] {
				j = mid - 1
			} else {
				i = mid + 1
			}
		}
	}
	return false
}
