/**
 * @Author: lihaoran
 * @Date: 2022/7/17 15:18
 */

package bubble_sort

import (
	"fmt"
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	s := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println(bubbleSort(s))
}
