/**
 * @Author: lihaoran04@baidu.com
 * @Date: 2020/7/30 10:06 下午
 */

package main

import (
	"fmt"
)

func main() {
	a := []int{9, 7, 8, 6, 7, 6, 5, 4, 3, 2, 1, 0}
	quickSort(a, 0, len(a)-1)
	//insertSort(a, 0, len(a)-1)
	fmt.Println(a)

}
