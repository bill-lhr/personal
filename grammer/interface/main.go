/**
 * @Author: lihaoran04@baidu.com
 * @Date: 2020/11/29 6:05 下午
 */

package main

import (
	"fmt"
	"reflect"
)

type stringer interface {
	string() string
}

type tester interface {
	stringer
	test()
}

type data struct {
	name string
}

type node struct {
	d interface {
		string() string
		test()
	}
}

func (d *data) test() {
	fmt.Println("test")
}
func (d *data) string() string {
	return d.name
}

func main() {
	var d data
	d.name = "lihaoran"

	//var t tester = d // 错误：data没有实现tester和stringer，因为data的receiver为指针，编译器会根据方法集判断是否实现了接口
	// 而对象对方法的调用不通过方法集判断，而是隐式的字段名（this）直接调用

	var t tester = &d
	t.test()
	fmt.Println(t.string())

	n := node{d: &data{name: "lihaoran"}}
	n.d.test()
	fmt.Println(n.d.string())
	fmt.Println(reflect.TypeOf(n).Name())
}

type Ner interface {
	a()
	b(int)
	c(string) string
}

type N int

func (N) a()               {}
func (*N) b(int)           {}
func (*N) c(string) string { return "" }

//func main() {
//	var n N
//	var ner Ner = &n
//	ner.a()
//}
