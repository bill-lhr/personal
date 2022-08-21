/**
 * @Author: lihaoran
 * @Date: 2022/8/21 15:30
 */

package hard

import "github.com/bill_lhr/personal/common"

/**
 * @Description: 寻找两个有序数组的中位数
	给定两个大小分别为 m 和 n 的正序（从小到大）数组nums1 和nums2。请你找出并返回这两个正序数组的中位数 。
	算法的时间复杂度应该为 O(log (m+n)) 。
	链接：https://leetcode.cn/problems/median-of-two-sorted-arrays
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	k := (len(nums1) + len(nums2)) / 2
	if (len(nums1)+len(nums2))%2 == 0 {
		return (float64(findKthMin(nums1, nums2, k)) + float64(findKthMin(nums1, nums2, k+1))) / 2
	}
	return float64(findKthMin(nums1, nums2, k+1))
}

// 递归
func findKthMin(nums1, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if len(nums2) == 0 {
		return nums1[k-1]
	}
	if k == 1 {
		return common.Min(nums1[0], nums2[0])
	}

	pivot := k / 2
	if len(nums1) < pivot || len(nums2) < pivot {
		pivot = common.Min(len(nums1), len(nums2))
	}
	if nums1[pivot-1] < nums2[pivot-1] {
		return findKthMin(nums1[pivot:], nums2, k-pivot)
	} else {
		return findKthMin(nums1, nums2[pivot:], k-pivot)
	}
}
