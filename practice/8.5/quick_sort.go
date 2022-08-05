package __5

func quickSort(n []int, i, j int) {
	if i >= j {
		return
	}
	if j-i+1 >= 3 {
		dealPivot(n, i, j)
		pivot, l, r := j-1, i+1, j-2
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
		n[l], n[pivot] = n[pivot], n[l]
		quickSort(n, i, l-1)
		quickSort(n, l+1, j)
	} else {
		if n[i] > n[j] {
			n[i], n[j] = n[j], n[i]
		}
	}
}

// 三数中值法
func dealPivot(n []int, l, r int) {
	mid := (l + r) / 2
	if n[l] > n[mid] {
		n[l], n[mid] = n[mid], n[l]
	}
	if n[mid] > n[r] {
		n[r], n[mid] = n[mid], n[r]
	}
	if n[l] > n[mid] {
		n[l], n[mid] = n[mid], n[l]
	}
	n[mid], n[r-1] = n[r-1], n[mid]
}
