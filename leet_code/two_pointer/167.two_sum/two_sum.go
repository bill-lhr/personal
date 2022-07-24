/**
 * @Author: lihaoran
 * @Date: 2022/7/24 11:09
 */

package two_sum

/*
 * @Description:
	给你一个下标从 1 开始的整数数组numbers ，该数组已按 非递减顺序排列 ，请你从数组中找出满足相加之和等于目标数target 的两个数。如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。
	以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。
	你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
	你所设计的解决方案必须只使用常量级的额外空间。
	链接：https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted
*/
func twoSum(numbers []int, target int) []int {
	res := make([]int, 0, 2)
	// 双指针相向移动
	low, high := 0, len(numbers)-1
	// 需要两个数所以不能相交
	for low < high {
		sum := numbers[low] + numbers[high]
		if sum == target {
			res = append(res, []int{low + 1, high + 1}...) // 题目要求下标从1开始
			break
		} else if sum < target {
			low++
		} else {
			high--
		}
	}
	return res
}
