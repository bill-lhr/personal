/**
 * @Author: lihaoran04@baidu.com
 * @Date: 2020/8/2 5:52 下午
 */

package main

import "fmt"

func main() {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	heapSort(a, len(a))
	//heapSort2(a, len(a))
	fmt.Println(a)
}

func heapSort(a []int, n int) {
	//一次大顶堆调整
	//从最后一个非叶子结点开始，从右向左、从下向上进行，时间复杂度O(logn)
	for i := n/2 - 1; i >= 0; i-- {
		sift(a, i, n)
	}

	//大顶堆调整后，堆顶元素为数组中最大元素，将其与最后一个元素交换，则最大元素就放在了最终的位置上
	//一次交换过后，将数组0~i的元素再次进行大顶堆调整，确定下一个元素的位置，执行n次后结束
	for i := n - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		sift(a, 0, i)
	}
}

//一次针对节点的堆调整，将i节点调整到堆中所在位置
func sift(a []int, i int, n int) {
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
