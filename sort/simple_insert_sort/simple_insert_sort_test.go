/**
 * @Author: lihaoran
 * @Date: 2022/7/17 15:11
 */

package simple_insert_sort

import (
	"fmt"
	"testing"
)

func TestSimpleInsertSort(t *testing.T) {
	s := []int{5, 4, 2, 3, 0, 1}
	fmt.Println(SimpleInsertSort(s))
}
