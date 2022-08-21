/**
 * @Author: lihaoran
 * @Date: 2022/7/17 15:55
 */

package quick_sort

import (
	"fmt"
	"testing"
)

func Test_quickSort(t *testing.T) {
	s := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	quickSort2(s, 0, len(s)-1)
	fmt.Println(s)
}
