package __8

func heapSort(nums []int) []int {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		sift(nums, i, n-1)
	}
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		sift(nums, 0, i-1)
	}
	return nums
}

func sift(nums []int, i, j int) {
	for 2*i+1 <= j {
		left, right := 2*i+1, 2*i+2
		maxChild := left
		if right <= j && nums[left] < nums[right] {
			maxChild = right
		}
		if nums[maxChild] > nums[i] {
			nums[i], nums[maxChild] = nums[maxChild], nums[i]
			i = maxChild
		} else {
			break
		}
	}
}
