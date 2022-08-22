package __22

import "github.com/bill_lhr/personal/common"

func pathSum437(root *common.TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var dfs func(node *common.TreeNode, target int) int
	dfs = func(node *common.TreeNode, target int) int {
		if node == nil {
			return 0
		}
		res := 0
		if node.Val == target {
			res += 1
		}
		res += dfs(node.Left, target-node.Val)
		res += dfs(node.Right, target-node.Val)
		return res
	}
	res := dfs(root, targetSum)
	res += dfs(root.Left, targetSum)
	res += dfs(root.Right, targetSum)
	return res
}
