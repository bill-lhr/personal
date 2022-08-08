package __8

func sort(n []int) []int {
	if len(n) > 0 {
		quickSort(n, 0, len(n)-1)
	}
	return n
}

func quickSort(n []int, i, j int) {
	if i >= j {
		return
	}

	if j-i >= 2 {
		dealPivot(n, i, j)
		pivot, left, right := j-1, i+1, j-2
		for left < right {
			for left < right && n[left] < n[pivot] {
				left++
			}
			for left < right && n[right] > n[pivot] {
				right--
			}
			if left < right {
				n[left], n[right] = n[right], n[left]
			}
		}
		n[pivot], n[left] = n[left], n[pivot]
		quickSort(n, i, left-1)
		quickSort(n, left+1, j)
	} else {
		if n[i] > n[j] {
			n[i], n[j] = n[j], n[i]
		}
	}
}

func dealPivot(n []int, i, j int) {
	mid := (i + j) / 2
	if n[i] > n[mid] {
		n[i], n[mid] = n[mid], n[i]
	}
	if n[mid] > n[j] {
		n[mid], n[j] = n[j], n[mid]
	}
	if n[i] > n[mid] {
		n[i], n[mid] = n[mid], n[i]
	}
	n[mid], n[j-1] = n[j-1], n[mid]
}
