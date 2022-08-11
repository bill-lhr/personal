package __11

import (
	"fmt"
	"testing"
)

func Test_quickSort(t *testing.T) {
	n := []int{7, 8, 9, 5, 6, 2, 3, 1, 0}
	quickSort(n, 0, len(n)-1)
	fmt.Println(n)
}
