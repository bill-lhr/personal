package sub_sets

// https://leetcode.cn/problems/subsets/
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	set := make([]int, 0)
	var backTracking func(int)

	backTracking = func(index int) {
		// 每一层都把结果输出
		tmp := make([]int, len(set))
		copy(tmp, set)
		res = append(res, tmp)
		// 横向遍历子集树
		for i := index; i < len(nums); i++ {
			set = append(set, nums[i])
			backTracking(i + 1)
			set = set[0 : len(set)-1]
		}
	}
	backTracking(0)
	return res
}
