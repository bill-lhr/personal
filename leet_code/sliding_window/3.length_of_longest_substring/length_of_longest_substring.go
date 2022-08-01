/**
 * @Author: lihaoran
 * @Date: 2022/7/25 23:35
 */

package length_of_longest_substring

/*
 * @Description:
	3. 无重复字符的最长子串
	给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
	链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters/
*/
func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	left, right := 0, 0
	m := make(map[byte]int)

	for right < len(s) {
		b := s[right]
		m[b]++
		right++

		for m[b] > 1 {
			m[s[left]]--
			left++
		}
		if right-left > maxLen {
			maxLen = right - left
		}
	}
	return maxLen
}
