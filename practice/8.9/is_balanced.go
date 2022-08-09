/**
 * @Author: lihaoran
 * @Date: 2022/8/9 22:37
 * @Description:
 */

package __9

import "github.com/bill_lhr/personal/common"

func isBalanced(root *common.TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(*common.TreeNode) int
	flag := true

	dfs = func(node *common.TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		diff := right - left
		if diff > 1 || diff < -1 {
			flag = false
		}
		if diff > 0 {
			return right + 1
		}
		return left + 1
	}
	return flag
}
