/**
 * @Author: lihaoran
 * @Date: 2022/7/24 23:04
 */

package min_window

/*
 * @Description:
	给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
	对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
	如果 s 中存在这样的子串，我们保证它是唯一的答案。
	链接：https://leetcode.cn/problems/minimum-window-substring
*/

func minWindow(s string, t string) string {
	left, right := 0, 0                                 // 滑动窗口指针
	has, need := make(map[byte]int), make(map[byte]int) // 分别表示窗口包含的字符和需要的字符
	valid := 0                                          // 表示每个总数全部覆盖的字符数量
	minStr := ""

	for i := range t {
		need[t[i]]++
	}

	for right < len(s) {
		b := s[right]
		if need[b] > 0 {
			has[b]++
			if has[b] == need[b] {
				valid++
			}
		}
		right++

		for valid == len(need) {
			// 更新结果要写在窗口缩小的循环里，因为缩小的过程中可能保持了结果的可行，但不能保证最小
			if minStr == "" || right-left < len(minStr) {
				minStr = s[left:right]
			}
			b = s[left]
			if need[b] > 0 {
				if has[b] == need[b] {
					valid--
				}
				has[b]--
			}
			left++
		}
	}
	return minStr
}
