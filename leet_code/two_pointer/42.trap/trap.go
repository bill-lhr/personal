/**
 * @Author: lihaoran
 * @Date: 2022/7/24 23:44
 */

package trap

import "github.com/bill_lhr/personal/common"

/*
 * @Description:
	给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/

func trap(height []int) int {
	// 1.双指针i、j分别从左右相向而行
	// 2.在指针移动的过程中维持maxLeft、maxRight分别表示当前左侧最大高度、右侧最大高度
	// 3.高度最大值小的那一侧的指针位置可以确定接雨量，且接雨量为该侧高度最大值减当前指针高度
	// 关键点：为什么高度最大值小的那一侧当前指针接雨量可以确定？
	if len(height) == 0 {
		return 0
	}
	sum := 0
	left, right, maxLeft, maxRight := 0, len(height)-1, 0, 0
	for left < right {
		maxLeft = common.Max(maxLeft, height[left])
		maxRight = common.Max(maxRight, height[right])
		if maxLeft < maxRight {
			sum += maxLeft - height[left]
			left++
		} else {
			sum += maxRight - height[right]
			right--
		}
	}
	return sum
}
