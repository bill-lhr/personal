package __5

import (
	"fmt"
	"testing"
)

func Test_quickSort(t *testing.T) {
	s := []int{3, 5, 4, 1, 2, 6}
	quickSort(s, 0, 5)
	fmt.Println(s)
}
