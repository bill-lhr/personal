package __11

func heapSort(n []int) []int {
	length := len(n)
	for i := length/2 - 1; i >= 0; i-- {
		sift(n, i, length-1)
	}

	for i := length - 1; i > 0; i-- {
		n[0], n[i] = n[i], n[0]
		sift(n, 0, i-1)
	}

	return n
}

func sift(n []int, i, j int) {
	for 2*i+1 <= j {
		left, right := 2*i+1, 2*i+2
		maxChild := left
		if right <= j && n[right] > n[left] {
			maxChild = right
		}
		if n[i] < n[maxChild] {
			n[i], n[maxChild] = n[maxChild], n[i]
			i = maxChild
		} else {
			break
		}
	}
}
