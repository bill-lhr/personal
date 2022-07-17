/**
 * @Author: lihaoran
 * @Description: 归并排序
 * @Date: 2022/7/17 18:00
 */

package merge_sort

func mergeSort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	mid := len(s) / 2
	left := mergeSort(s[0:mid])
	right := mergeSort(s[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	s := make([]int, 0)
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			s = append(s, left[i])
			i++
		} else {
			s = append(s, right[j])
			j++
		}
	}
	for i < len(left) {
		s = append(s, left[i])
	}
	for j < len(right) {
		s = append(s, right[j])
	}
	return s
}
