/**
 * @Author: lihaoran
 * @Date: 2022/8/27 15:24
 */

package binary_search

/**
 * @Description: 整数数组 nums 按升序排列，数组中的值 互不相同 。
	给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
	链接: https://leetcode.cn/problems/search-in-rotated-sorted-array/
*/
func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < nums[j] {
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
	return -1
}
