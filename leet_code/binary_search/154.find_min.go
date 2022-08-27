/**
 * @Author: lihaoran
 * @Date: 2022/8/27 15:09
 */

package binary_search

/**
 * @Description: 已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。
	给你一个可能存在 重复 元素值的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
	链接: https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/
*/
func findMin(nums []int) int {
	i, j := 0, len(nums)-1
	for i < j {
		mid := (i + j) / 2
		if nums[mid] > nums[j] {
			i = mid + 1
		} else if nums[mid] < nums[j] {
			j = mid
		} else {
			j = j - 1
		}
	}
	return nums[i]
}
