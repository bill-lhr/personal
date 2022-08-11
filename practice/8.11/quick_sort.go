package __11

func quickSort(n []int, i, j int) {
	if i >= j {
		return
	}

	if j-i+1 >= 3 {
		dealPivot(n, i, j)
		left, right, pivot := i+1, j-2, j-1
		for left < right {
			for left < right && n[left] < n[pivot] {
				left++
			}
			for left < right && n[right] > n[pivot] {
				right--
			}
			if left < right {
				swap(n, left, right)
			}
		}
		swap(n, left, pivot)
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
		swap(n, i, mid)
	}
	if n[mid] > n[j] {
		swap(n, mid, j)
	}
	if n[i] > n[mid] {
		swap(n, i, mid)
	}
	swap(n, mid, j-1)
}

func swap(n []int, i, j int) {
	n[i], n[j] = n[j], n[i]
}
