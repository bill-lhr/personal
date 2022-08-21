/**
 * @Author: lihaoran
 * @Date: 2022/8/21 20:03
 */

package hard

/**
 * @Description: 76.最小覆盖子串
	给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
	链接: https://leetcode.cn/problems/minimum-window-substring/
*/
func minWindow(s string, t string) string {
	need, has := make(map[byte]int), make(map[byte]int)
	valid := 0
	left, right := 0, 0
	res := ""
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
			if res == "" || right-left < len(res) {
				res = s[left:right]
			}
			b = s[left]
			if has[b] > 0 {
				if has[b] == need[b] {
					valid--
				}
				has[b]--
			}
			left++
		}
	}
	return res
}
