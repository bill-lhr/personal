package check_inclusion

import (
	"fmt"
	"testing"
)

func Test_checkInclusion(t *testing.T) {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
}
