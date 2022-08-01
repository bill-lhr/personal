/**
 * @Author: lihaoran
 * @Date: 2022/8/1 22:26
 * @Description:
 */

package merge

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	a := make([]int, 7)
	a[0], a[1], a[2], a[3] = 1, 2, 3, 4
	merge(a, 4, []int{1, 2, 3}, 3)
	fmt.Println(a)
}
