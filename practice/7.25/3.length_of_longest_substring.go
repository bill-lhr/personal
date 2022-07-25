/**
 * @Author: lihaoran
 * @Date: 2022/7/25 23:14
 * @Description:
 */

package _725

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	maxLen := 0
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
