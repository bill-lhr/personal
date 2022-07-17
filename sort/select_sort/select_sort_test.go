/**
 * @Author: lihaoran
 * @Date: 2022/7/17 17:47
 */

package select_sort

import (
	"fmt"
	"testing"
)

func Test_selectSort(t *testing.T) {
	s := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 10, 11, 12}
	selectSort(s)
	fmt.Println(s)
}
