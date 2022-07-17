/**
 * @Author: lihaoran
 * @Date: 2022/7/17 16:31
 */

package heap_sort

func heapSort(s []int) []int {
	//一次大顶堆调整
	//从最后一个非叶子结点开始，从右向左、从下向上进行，时间复杂度O(logn)
	for i := len(s)/2 - 1; i >= 0; i-- {
		sift(s, i, len(s)-1)
	}

	//大顶堆调整后，堆顶元素为数组中最大元素，将其与最后一个元素交换，则最大元素就放在了最终的位置上
	//一次交换过后，将数组0~i的元素再次进行大顶堆调整，确定下一个元素的位置，执行n次后结束
	for i := len(s) - 1; i > 0; i-- {
		s[0], s[i] = s[i], s[0]
		sift(s, 0, i-1)
	}
	return s
}

// 大顶堆调整 i表示被调整的节点索引 n表示堆结束的位置
func sift(s []int, i, n int) {
	for 2*i+1 <= n {
		left, right := 2*i+1, 2*i+2
		maxChild := left
		if right <= n && s[right] > s[left] {
			maxChild = right
		}
		if s[maxChild] > s[i] {
			s[i], s[maxChild] = s[maxChild], s[i]
			i = maxChild
		} else {
			break
		}
	}
}
