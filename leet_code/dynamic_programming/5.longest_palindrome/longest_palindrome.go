package longest_palindrome

/**
 * @Description:
	最长回文子串 https://leetcode.cn/problems/longest-palindromic-substring/
	给你一个字符串 s，找到 s 中最长的回文子串。

	1. 定义: dp[i][j] 表示从i到j是否为回文
	2. 状态转移: i==j, dp[i][j]=true; i+1==j, dp[i][j]=true; s[i]==s[j], dp[i][j]=dp[i+1][j-1]; s[i]!=s[j], dp[i][j]=false
	3. 初始化: dp[i][i]=true
	4. 遍历顺序: 此题关键点，需要按长度进行遍历，每次固定一个长度然后确定开始结束的i,j
	5. 结果: dp[i][j]==true时记录 j-i+1 的最大值
*/
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	maxLen, left := 1, 0
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	for length := 2; length <= n; length++ {
		for i := 0; i < n-length+1; i++ { // 注意i的最大值
			j := i + length - 1
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if i == j || i+1 == j {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
				if dp[i][j] && j-i+1 > maxLen {
					maxLen, left = j-i+1, i
				}
			}
		}
	}
	return s[left : left+maxLen]
}
