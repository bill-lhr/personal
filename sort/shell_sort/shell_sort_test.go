/**
 * @Author: lihaoran
 * @Date: 2022/7/17 15:12
 */

package shell_sort

import (
	"fmt"
	"testing"
)

func Test_shellSort(t *testing.T) {
	s := []int{5, 4, 2, 3, 0, 1}
	fmt.Println(shellSort(s))
}
