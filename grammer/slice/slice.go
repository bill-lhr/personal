/**
 * @Author: lihaoran04@baidu.com
 * @Date: 2020/8/30 9:53 下午
 */

package main

import "fmt"

func main() {
	// slice 扩容实验
	s := make([]int, 0, 3)
	fmt.Printf("slice: %v, pointer: %p\n", s, s)
	s = append(s, 1, 2)
	fmt.Printf("slice: %v, pointer: %p\n", s, s)
	sliceAppend(s)
	fmt.Printf("slice: %v, pointer: %p\n", s, s)
	str := "12345"
	fmt.Println(string(str[0]))
}

func sliceAppend(s []int) {
	s = append(s, 1)
	s[0] = 2
	fmt.Printf("slice in func: %v, pointer: %p\n", s, s)
}
