package __22

import "github.com/bill_lhr/personal/common"

func pathSum(root *common.TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(node *common.TreeNode, target int)
	dfs = func(node *common.TreeNode, target int) {
		if node == nil {
			return
		}
		path = append(path, root.Val)
		if root.Val == target && root.Left == nil && root.Right == nil {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		dfs(root.Left, target-root.Val)
		dfs(root.Right, target-root.Val)
		path = path[:len(path)-1]
	}
	dfs(root, targetSum)
	return res
}
