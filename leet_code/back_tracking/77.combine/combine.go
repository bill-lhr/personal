package combine

// https://leetcode.cn/problems/combinations/
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	list := make([]int, 0)
	var backTracking func(int)
	backTracking = func(index int) {
		// 递归退出条件
		if len(list) == k {
			// 记录结果
			tmp := make([]int, len(list))
			copy(tmp, list)
			res = append(res, tmp)
			return
		}

		// 子集树每层的遍历
		for i := index; i <= n; i++ {
			// 节点选择
			list = append(list, i)
			backTracking(i + 1)
			// 回溯
			list = list[0 : len(list)-1]
		}
	}

	backTracking(1)
	return res
}
