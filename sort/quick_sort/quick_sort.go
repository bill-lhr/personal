/**
 * @Author: lihaoran
 * @Date: 2022/7/17 15:51
 */

package quick_sort

func quickSort(s []int, left, right int) {
	if left >= right {
		return
	}

	pivot, i, j := left, left, right
	for i < j {
		for i < j && s[i] < s[pivot] {
			i++
		}
		for i < j && s[j] > s[pivot] {
			j--
		}
		if i < j {
			s[i], s[j] = s[j], s[i]
		}
	}
	s[i], s[pivot] = s[pivot], s[i]

	quickSort(s, left, i-1)
	quickSort(s, i+1, right)
}
