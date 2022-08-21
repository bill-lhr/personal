/**
 * @Author: lihaoran
 * @Date: 2022/8/21 17:33
 */

package hard

/**
 * @Description: 最长有效括号
	给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
	链接: https://leetcode.cn/problems/longest-valid-parentheses/
*/

// 方法一： 栈
func longestValidParentheses(s string) int {
	maxLen := 0
	stack := make([]int, 0)
	stack = append(stack, -1)
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) == 1 {
				stack[0] = i
			} else {
				stack = stack[:len(stack)-1]
				length := i - stack[len(stack)-1]
				if length > maxLen {
					maxLen = length
				}
			}
		}
	}
	return maxLen
}

// 方法二：动态规划
func longestValidParentheses2(s string) int {
	maxLen := 0
	n := len(s)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		if s[i] == '(' {
			continue
		}
		if s[i-1] == '(' {
			if i > 1 {
				dp[i] = dp[i-2] + 2
			} else {
				dp[i] = 2
			}
		} else { // s[i-1] == ')'
			if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1]-2 >= 0 {
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}
