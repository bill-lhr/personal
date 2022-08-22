package __17

func quickSort(n []int, i, j int) {
	if i >= j {
		return
	}

	if j-i >= 2 {
		dealPivot(n, i, j)
		l, r, pivot := i+1, j-2, j-1
		for l < r {
			for l < r && n[l] <= n[pivot] {
				l++
			}
			for l < r && n[r] >= n[pivot] {
				r--
			}
			if l < r {
				n[l], n[r] = n[r], n[l]
			}
		}
		n[l], n[pivot] = n[pivot], n[l]
		quickSort(n, i, l-1)
		quickSort(n, l+1, j)
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
