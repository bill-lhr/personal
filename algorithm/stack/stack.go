/**
 * @Author: lihaoran04@baidu.com
 * @Date: 2020/8/23 11:42 下午
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// string本质是一个 []byte, byte即 uint8
	s := "{}()"
	fmt.Println(len(s))
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	var st stack
	st.data = make([]string, 0)
	for _, b := range s {
		// 当使用range遍历string时，返回的是unicode，类型是rune --- int32
		fmt.Println(reflect.TypeOf(b))
		a := string(b)
		switch a {
		case "(", "{", "[":
			st.push(a)
		case ")", "}", "]":
			aa := st.pop()
			if (a == ")" && aa != "(") || (a == "}" && aa != "{") || (a == "]" && aa != "[") {
				return false
			}
		}
	}
	//for i := 0; i < len(s); i++ {
	//	fmt.Println(s[i])
	//	fmt.Printf("%s\n", string(s[i]))
	//}
	return st.isEmpty()
}

type stack struct {
	data []string
}

func (s *stack) pop() string {
	if s.isEmpty() {
		return ""
	}
	r := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return r
}

func (s *stack) push(a string) {
	s.data = append(s.data, a)
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}
