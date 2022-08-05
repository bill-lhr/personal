/**
 * @Author: lihaoran
 * @Date: 2022/8/5 22:22
 * @Description:
 */

package next_permutation

import (
	"fmt"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	nums := []int{1, 2, 3, 6, 5, 4}
	nextPermutation(nums)
	fmt.Println(nums)
}
