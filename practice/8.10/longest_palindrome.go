/**
 * @Author: lihaoran
 * @Date: 2022/8/10 22:24
 * @Description:
 */

package __10

func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	res := s[0:1]
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	// 固定长度 遍历起点
	for length := 2; length <= n; length++ {
		for i := 0; i+length <= n; i++ {
			j := i + length - 1
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j == i+1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}
	return res
}
