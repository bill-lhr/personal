/**
 * @Author: lihaoran
 * @Date: 2022/7/25 22:08
 * @Description:
 */

package _725

func quickSort(nums []int) []int {
	var quick func([]int, int, int)
	quick = func(n []int, i, j int) {
		if i >= j {
			return
		}
		pivot, l, r := i, i, j
		for l < r {
			for l < r && n[l] < n[pivot] {
				l++
			}
			for l < r && n[r] > n[pivot] {
				r--
			}
			if l < r {
				n[l], n[r] = n[r], n[l]
			}
		}
		n[pivot], n[l] = n[l], n[pivot]

		quick(n, i, l-1)
		quick(n, l+1, j)
	}
	quick(nums, 0, len(nums)-1)

	return nums
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2 // 注意，这里必须是 len(nums)/2 否则只有两个元素时 mid == 0
	left := mergeSort(nums[0:mid])
	right := mergeSort(nums[mid:])

	return merge(left, right)
}

func merge(a, b []int) []int {
	lenA, lenB := len(a), len(b)
	res := make([]int, 0, lenA+lenB)
	i, j := 0, 0

	for i < lenA && j < lenB {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}

	for i < lenA {
		res = append(res, a[i])
		i++
	}
	for j < lenB {
		res = append(res, b[j])
		j++
	}

	return res
}

func insertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return nums
}

func heapSort(nums []int) []int {
	for i := len(nums)/2 + 1; i >= 0; i-- {
		sift(nums, i, len(nums)-1)
	}

	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		sift(nums, 0, i-1)
	}

	return nums
}

func sift(nums []int, i, n int) {
	for 2*i+1 <= n {
		left, right := 2*i+1, 2*i+2
		maxChild := left
		if right <= n && nums[right] > nums[left] {
			maxChild = right
		}
		if nums[maxChild] > nums[i] {
			nums[maxChild], nums[i] = nums[i], nums[maxChild]
			i = maxChild
		} else {
			break
		}
	}
}
