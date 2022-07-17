/**
 * @Author: lihaoran
 * @Date: 2022/7/17 15:10
 */

package bubble_sort

func bubbleSort(s []int) []int {
	for i := len(s) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}
