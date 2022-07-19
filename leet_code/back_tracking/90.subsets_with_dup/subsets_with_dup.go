package subsets_with_dup

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	list := make([]int, 0)
	var backTracking func(int)
	// 排序是为了保证元素相对顺序不变防止重复选择
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	backTracking = func(index int) {
		tmp := make([]int, len(list))
		copy(tmp, list)
		res = append(res, tmp)

		for i := index; i < len(nums); i++ {
			num := nums[i]
			// 剪枝
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			list = append(list, num)
			backTracking(i + 1)
			list = list[0 : len(list)-1]
		}
	}

	backTracking(0)
	return res
}
