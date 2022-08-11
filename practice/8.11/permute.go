package __11

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	list := make([]int, 0, len(nums))
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
			if used[nums[i]] {
				continue
			}
			list = append(list, nums[i])
			used[nums[i]] = true
			backTracking()
			list = list[:len(list)-1]
			used[nums[i]] = false
		}
	}
	backTracking()
	return res
}
