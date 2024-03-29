/**
 * @Author: lihaoran
 * @Date: 2022/7/24 11:25
 */

package remove_duplicates

/*
 * @Description:
	给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。
	由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果。
	将最终结果插入nums 的前 k 个位置后返回 k 。
	链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-array
*/
func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
		fast++
	}
	return slow + 1
}
