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
	s := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	quickSort(s, 0, len(s)-1)
	fmt.Println(s)
}
