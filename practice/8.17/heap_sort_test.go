package __17

import (
	"fmt"
	"testing"
)

func Test_heapSort(t *testing.T) {
	n := []int{7, 8, 9, 5, 6, 2, 3, 1, 0}
	fmt.Println(heapSort(n))
}
