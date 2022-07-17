/**
 * @Author: lihaoran
 * @Description:
 * @Date: 2022/7/17 18:06
 */

package merge_sort

import (
	"fmt"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	s := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 10, 11, 12}

	fmt.Println(mergeSort(s))
}
