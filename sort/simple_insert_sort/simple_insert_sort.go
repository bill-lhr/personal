/**
 * @Author: lihaoran
 * @Date: 2022/7/17 11:21
 */

package simple_insert_sort

func SimpleInsertSort(s []int) []int {
	for i := 0; i < len(s); i++ {
		tmp := s[i]
		j := i - 1
		for ; j >= 0; j-- {
			if s[j] > tmp {
				s[j+1] = s[j]
			} else {
				break
			}
		}
		s[j+1] = tmp
	}
	return s
}
