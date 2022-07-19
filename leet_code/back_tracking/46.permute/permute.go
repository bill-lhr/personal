package permute

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	list := make([]int, 0)
	used := make(map[int]bool)
	var backTracking func()

	backTracking = func() {
		if len(list) == len(nums) {
			tmp := make([]int, len(list))
			copy(tmp, list)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			list = append(list, nums[i])
			used[i] = true
			backTracking()
			used[i] = false
			list = list[0 : len(list)-1]
		}
	}

	backTracking()
	return res
}
