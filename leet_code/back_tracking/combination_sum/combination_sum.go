package combination_sum

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	list := make([]int, 0)
	sum := 0
	var backTracking func(int)

	backTracking = func(index int) {
		if sum == target {
			tmp := make([]int, len(list))
			copy(tmp, list)
			res = append(res, tmp)
			return
		}
		if sum > target {
			return
		}

		for i := index; i < len(candidates); i++ {
			num := candidates[i]
			sum += num
			list = append(list, num)
			backTracking(i)
			sum -= num
			list = list[0 : len(list)-1]
		}
	}

	backTracking(0)

	return res
}
