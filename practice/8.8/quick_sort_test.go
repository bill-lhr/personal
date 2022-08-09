package __8

import (
	"fmt"
	"testing"
)

func Test_sort(t *testing.T) {
	fmt.Println(sort([]int{7, 8, 9, 5, 6, 2, 3, 1, 0}))
	fmt.Println(sort([]int{8, 7}))
	fmt.Println(sort([]int{1}))
	fmt.Println(heapSort([]int{7, 8, 9, 5, 6, 2, 3, 1, 0}))
}
