/**
 * @Author: lihaoran
 * @Date: 2022/7/17 14:41
 */

package shell_sort

import "fmt"

func shellSort(s []int) []int {
	length := len(s)
	// 间隔gap从length/2开始，每次缩小为gap/2
	for gap := length / 2; gap > 0; gap /= 2 {
		// 通过gap分组，每组内进行插入排序
		for i := gap; i < length; i++ {
			tmp, j := s[i], i
			for ; j >= gap; j -= gap {
				if s[j-gap] > tmp {
					s[j] = s[j-gap]
				} else {
					break
				}
			}
			s[j] = tmp
		}
	}
	return s
}

func main() {
	s := []int{5, 4, 2, 3, 0, 1}
	fmt.Println(shellSort(s))
}
