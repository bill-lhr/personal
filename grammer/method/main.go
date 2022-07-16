/**
 * @Author: lihaoran04@baidu.com
 * @Date: 2020/11/29 5:05 下午
 */

package main

import (
	"fmt"
	"reflect"
)

type S struct{}

type T struct {
	S
}

func (S) SVal()  {}
func (*S) SPtr() {}
func (T) TVal()  {}
func (*T) TPtr() {}

func methodSet(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println(t.NumMethod())

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

func main() {
	var t T

	methodSet(t)
	println("---------------")
	methodSet(&t)
}
