package reverse_pairs

/**
 * @Description:
	在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
	链接: https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
*/
func reversePairs(nums []int) int {
	cnt := 0
	var mergeSort func([]int) []int
	var merge func([]int, []int) []int

	mergeSort = func(n []int) []int {
		if len(n) < 2 {
			return n
		}
		mid := len(n) / 2
		left := mergeSort(n[:mid])
		right := mergeSort(n[mid:])
		return merge(left, right)
	}

	merge = func(n1 []int, n2 []int) []int {
		left, right := 0, 0
		res := make([]int, 0, len(n1)+len(n2))
		for left < len(n1) && right < len(n2) {
			if n1[left] < n2[right] {
				res = append(res, n1[left])
				left++
			} else {
				cnt += len(n1) - left
				res = append(res, n2[right])
				right++
			}
		}
		if left < len(n1) {
			res = append(res, n1[left:]...)
		}
		if right < len(n2) {
			res = append(res, n2[right:]...)
		}
		return res
	}

	mergeSort(nums)
	return cnt
}
