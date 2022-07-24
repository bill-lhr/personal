/**
 * @Author: lihaoran
 * @Date: 2022/7/24 15:32
 */

package find_content_children

import "sort"

/*
 * @Description:
	假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。
	对每个孩子 i，都有一个胃口值g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，都有一个尺寸 s[j]。
	如果 s[j]>= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。
	链接：https://leetcode.cn/problems/assign-cookies
*/

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	gIdx, sIdx, cnt := 0, 0, 0
	for gIdx < len(g) && sIdx < len(s) {
		if g[gIdx] <= s[sIdx] {
			cnt++
			gIdx++
			sIdx++
		} else {
			sIdx++
		}
	}

	return cnt
}
