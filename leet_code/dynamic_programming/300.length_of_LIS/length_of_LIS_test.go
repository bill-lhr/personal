/**
 * @Author: lihaoran
 * @Description:
 * @Date: 2022/7/20 23:50
 */

package length_of_LIS

import (
	"fmt"
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func Test_lengthOfLIS2(t *testing.T) {
	fmt.Println(lengthOfLIS2([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
}
