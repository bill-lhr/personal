package build_tree

import "github.com/bill_lhr/personal/common"

func buildTree(preorder []int, inorder []int) *common.TreeNode {
	if len(preorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}

	root := &common.TreeNode{Val: preorder[0]}
	idx := 0
	for ; idx < len(inorder); idx++ {
		if inorder[idx] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:idx+1], inorder[0:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])

	return root
}
