package __18

/**
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
*/
func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	maxLen := 0
	has := make(map[byte]int)
	for right < len(s) {
		b := s[right]
		has[b]++
		right++

		for has[b] > 1 {
			has[s[left]]--
			left++
		}
		if right-left > maxLen {
			maxLen = right - left
		}
	}
	return maxLen
}
