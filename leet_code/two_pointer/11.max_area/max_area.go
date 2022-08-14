/**
 * @Author: lihaoran
 * @Date: 2022/8/14 18:08
 * @Description:
 */

package max_area

import "github.com/bill_lhr/personal/common"

func maxArea(height []int) int {
	res, left, right := 0, 0, len(height)-1
	for left < right {
		area := (right - left) * common.Min(height[left], height[right])
		if area > res {
			res = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}
