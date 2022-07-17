/**
 * @Author: lihaoran
 * @Date: 2022/7/17 16:46
 */

package heap_sort

import (
	"fmt"
	"testing"
)

func Test_heapSort(t *testing.T) {
	s := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 10, 11, 12}
	heapSort(s)
	fmt.Println(s)
}
