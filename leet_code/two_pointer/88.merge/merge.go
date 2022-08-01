/**
 * @Author: lihaoran
 * @Date: 2022/8/1 22:17
 */

package merge

/**
 * @Description: 给你两个按 非递减顺序 排列的整数数组nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
	请你合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
	链接：https://leetcode.cn/problems/merge-sorted-array
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
