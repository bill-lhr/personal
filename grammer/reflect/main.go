/**
 * @Author: lihaoran04@baidu.com
 * @Date: 2020/11/29 9:26 下午
 */

package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

func main() {
	var m manager
	t := reflect.TypeOf(&m)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type, f.Offset)
		if f.Anonymous {
			for j := 0; j < f.Type.NumField(); j++ {
				af := t.Field(i)
				fmt.Println(af.Name, af.Type, af.Offset)
			}
		}
	}
}
