/**
 * @Author: lihaoran04@baidu.com
 * @Date: 2020/8/2 7:16 下午
 */

package main

func heapSort2(a []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		sift2(a, i, n)
	}

	for i := n - 1; i > 0; i-- {
		a[i], a[0] = a[0], a[i]
		sift2(a, 0, i-1)
	}
}

func sift2(a []int, i int, n int) {
	for 2*i+1 < n {
		left, right := 2*i+1, 2*i+2
		maxChild := left
		if right < n && a[left] < a[right] {
			maxChild = right
		}
		if a[i] < a[maxChild] {
			a[i], a[maxChild] = a[maxChild], a[i]
			i = maxChild
		} else {
			break
		}
	}
}
