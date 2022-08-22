package __22

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	list := make([]int, 0)
	var backTracking func(int)
	backTracking = func(idx int) {
		tmp := make([]int, len(list))
		copy(tmp, list)
		res = append(res, tmp)

		for i := idx; i < len(nums); i++ {
			list = append(list, nums[i])
			backTracking(i + 1)
			list = list[:len(list)-1]
		}
	}
	backTracking(0)
	return res
}
