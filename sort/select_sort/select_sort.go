/**
 * @Author: lihaoran
 * @Date: 2022/7/17 17:42
 */

package select_sort

func selectSort(s []int) {
	for i := 0; i < len(s); i++ {
		minIndex := i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[i] {
				minIndex = j
			}
		}
		s[i], s[minIndex] = s[minIndex], s[i]
	}
}
