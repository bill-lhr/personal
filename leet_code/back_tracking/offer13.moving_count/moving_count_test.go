/**
 * @Author: lihaoran
 * @Description:
 * @Date: 2022/7/21 22:45
 */

package moving_count

import (
	"fmt"
	"testing"
)

func Test_check(t *testing.T) {
	fmt.Println(check(35, 38, 19))
}

func Test_movingCount(t *testing.T) {
	fmt.Println(movingCount(7, 2, 3))
}
