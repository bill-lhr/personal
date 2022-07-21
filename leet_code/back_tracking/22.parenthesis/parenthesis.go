package parenthesis

// https://leetcode.cn/problems/generate-parentheses/
// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
// eg: n == 2 res: (())，()()
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	list := make([]byte, 0) // 用来存放一个解
	l, r := 0, 0            // 表示左右括号个数
	var backTracking func()

	backTracking = func() {
		// 递归退出 记录结果
		if l == n && r == n {
			s := string(list)
			res = append(res, s)
			return
		}

		for _, item := range []byte{'(', ')'} {
			// 选择
			if item == '(' && l < n {
				list = append(list, item)
				l++
				backTracking()
				l--
				list = list[0 : len(list)-1]
			}
			if item == ')' && r < l {
				list = append(list, item)
				r++
				backTracking()
				r--
				list = list[0 : len(list)-1]
			}
		}
	}

	backTracking()
	return res
}
