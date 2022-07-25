/**
 * @Author: lihaoran
 * @Date: 2022/7/25 22:18
 * @Description:
 */

package _725

import (
	"fmt"
	"testing"
)

func Test_quickSort(t *testing.T) {
	fmt.Println(quickSort([]int{4, 3, 5, 1, 2}))
}

func Test_mergeSort(t *testing.T) {
	fmt.Println(mergeSort([]int{4, 3, 5, 2, 1}))
}

func Test_insertSort(t *testing.T) {
	fmt.Println(insertSort([]int{4, 3, 5, 2, 1}))
}

func Test_heapSort(t *testing.T) {
	fmt.Println(heapSort([]int{4, 3, 5, 2, 1}))
}
