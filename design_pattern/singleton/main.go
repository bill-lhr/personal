/**
 * @Author: lihaoran
 * @Date: 2022/6/30 23:34
 */

package main

import (
	"fmt"
	"sync"
)

var (
	singleton *Singleton
	once      = &sync.Once{}
)

func getSingleton() *Singleton {
	once.Do(func() {
		singleton = &Singleton{
			field: "hhhh",
		}
	})
	return singleton
}

type Singleton struct {
	field interface{}
}

func main() {
	fmt.Println(getSingleton().field)
}
