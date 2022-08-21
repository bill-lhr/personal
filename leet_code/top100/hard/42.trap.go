/**
 * @Author: lihaoran
 * @Date: 2022/8/21 19:41
 */

package hard

import "github.com/bill_lhr/personal/common"

/**
 * @Description: 接雨水
	给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
	链接: https://leetcode.cn/problems/trapping-rain-water/
*/
func trap(height []int) int {
	res := 0
	left, right := 0, len(height)-1
	maxLeft, maxRight := 0, 0
	for left <= right {
		maxLeft = common.Max(maxLeft, height[left])
		maxRight = common.Max(maxRight, height[right])
		if maxLeft < maxRight {
			res += maxLeft - height[left]
		} else {
			res += maxRight - height[right]
		}
	}
	return res
}
