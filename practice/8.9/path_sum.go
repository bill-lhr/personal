/**
 * @Author: lihaoran
 * @Date: 2022/8/9 23:29
 * @Description:
 */

package __9

import "github.com/bill_lhr/personal/common"

// 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
func pathSum(root *common.TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var dfs func(node *common.TreeNode, target int) int
	dfs = func(node *common.TreeNode, target int) int {
		if node == nil {
			return 0
		}
		cnt := 0
		if node.Val == target {
			cnt += 1
		}
		cnt += dfs(node.Left, target-node.Val)
		cnt += dfs(node.Right, target-node.Val)
		return cnt
	}

	res := dfs(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}
