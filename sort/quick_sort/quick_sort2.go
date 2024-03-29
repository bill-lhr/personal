/**
 * @Author: lihaoran04@baidu.com
 * @Date: 2021/5/31 下午6:34
 */

package quick_sort

func quickSort2(a []int, low, high int) {
	if low >= high {
		return
	}

	if high-low+1 >= 3 {
		dealPivot(a, low, high)
		i, j := low+1, high-2
		pivot := high - 1
		for i < j {
			for i < j && a[i] <= a[pivot] {
				i++
			}
			for i < j && a[j] >= a[pivot] {
				j--
			}
			if i < j {
				a[i], a[j] = a[j], a[i]
			}
		}

		if i < pivot && a[i] > a[pivot] {
			a[i], a[pivot] = a[pivot], a[i]
		} else {
			i = pivot
		}
		quickSort2(a, low, i-1)
		quickSort2(a, i+1, high)
	} else {
		if a[low] > a[high] {
			a[low], a[high] = a[high], a[low]
		}
	}
}

func dealPivot(a []int, low, high int) {
	mid := (low + high) / 2
	if a[low] > a[mid] {
		a[low], a[mid] = a[mid], a[low]
	}
	if a[low] > a[high] {
		a[low], a[high] = a[high], a[low]
	}
	if a[mid] > a[high] {
		a[mid], a[high] = a[high], a[mid]
	}
	a[mid], a[high-1] = a[high-1], a[mid]
}

func insertSort(a []int) {
	for i := 1; i < len(a); i++ {
		tmp, j := a[i], i-1
		for ; j >= 0 && a[j] > tmp; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = tmp
	}
}
