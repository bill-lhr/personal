/**
 * @Author: lihaoran
 * @Date: 2022/7/16 19:06
 */

package check_inclusion

// https://leetcode.cn/problems/permutation-in-string/
func checkInclusion(s1 string, s2 string) bool {
	needs, has := make(map[byte]int), make(map[byte]int)
	left, right, valid := 0, 0, 0

	for i := range s1 {
		needs[s1[i]] += 1
	}

	for right < len(s2) {
		b := s2[right]
		if needs[b] > 0 {
			has[b]++
			if has[b] == needs[b] {
				valid++
			}
		}
		right++

		for valid == len(needs) {
			if right-left == len(s1) {
				return true
			}
			b := s2[left]
			if needs[b] > 0 {
				if has[b] == needs[b] {
					valid--
				}
				has[b]--
			}
			left++
		}
	}
	return false
}
