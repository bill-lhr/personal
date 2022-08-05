/**
 * @Author: lihaoran
 * @Date: 2022/8/5 23:16
 */

package max_product

import "github.com/bill_lhr/personal/common"

/**
 * @Description: 给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
	链接: https://leetcode.cn/problems/maximum-product-subarray/

	需要两个dp数组，分别表示以当前元素结尾的最大/最小的乘积，因为负数的存在最大和最小有可能互换
*/
func maxProduct(nums []int) int {
	res := nums[0]
	dpMax := make([]int, len(nums))
	dpMin := make([]int, len(nums))
	dpMax[0], dpMin[0] = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		dpMax[i] = common.Max(dpMax[i-1]*nums[i], common.Max(dpMin[i-1]*nums[i], nums[i]))
		dpMin[i] = common.Min(dpMax[i-1]*nums[i], common.Min(dpMin[i-1]*nums[i], nums[i]))
		if dpMax[i] > res {
			res = dpMax[i]
		}
	}
	return res
}
