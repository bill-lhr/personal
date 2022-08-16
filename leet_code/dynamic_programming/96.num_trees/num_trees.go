/**
 * @Author: lihaoran
 * @Date: 2022/8/16 22:55
 */

package num_trees

/**
 * @Description:
	给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
	链接: https://leetcode.cn/problems/unique-binary-search-trees/
*/
func numTrees(n int) int {
	// dp[i]表示i个连续整数能组成的不同BST个数
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
