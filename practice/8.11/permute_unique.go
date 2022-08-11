package __11

// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	list := make([]int, 0)
	used := make(map[int]int)
	has := make(map[int]int)
	for _, n := range nums {
		has[n]++
	}

	var backTracking func()
	backTracking = func() {
		if len(list) == len(nums) {
			tmp := make([]int, len(list))
			copy(tmp, list)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[nums[i]] == has[nums[i]] {
				continue
			}
			n := nums[i]
			used[n]++
			list = append(list, n)
			backTracking()
			used[n]--
			list = list[:len(list)-1]
		}
	}
	backTracking()
	return res
}
