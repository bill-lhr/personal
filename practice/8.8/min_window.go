package __8

// 最小覆盖子串
// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
func minWindow(s string, t string) string {
	has, need := make(map[byte]int), make(map[byte]int)
	left, right := 0, 0
	valid := 0
	res := ""
	for i := range t {
		need[t[i]]++
	}
	for right < len(s) {
		has[s[right]]++
		if has[s[right]] == need[s[right]] {
			valid++
		}
		right++

		for valid == len(need) {
			if len(res) == 0 || right-left < len(res) {
				res = s[left:right]
			}
			if has[s[left]] == need[s[left]] {
				valid--
			}
			has[s[left]]--
			left++
		}
	}
	return res
}
