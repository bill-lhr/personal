/**
 * @Author: lihaoran04@baidu.com
 * @Date: 2020/8/2 5:40 下午
 */

package main

import "fmt"

func main() {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	insertSort(a)
	fmt.Println(a)
}

func insertSort(a []int) {
	for i := 1; i < len(a); i++ {
		tmp, j := a[i], i-1
		for ; j >= 0 && a[j] > tmp; j-- {
			a[j+1] = a[j]
		}
		// 注意：循环结束后j执行了减1，所以最终在j+1的位置插入
		a[j+1] = tmp
	}
}
