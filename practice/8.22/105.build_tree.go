package __22

import "github.com/bill_lhr/personal/common"

func buildTree(preorder []int, inorder []int) *common.TreeNode {
	if len(preorder) != len(inorder) || len(preorder) == 0 {
		return nil
	}
	root := &common.TreeNode{Val: preorder[0]}
	idx := 0
	for ; idx < len(inorder); idx++ {
		if inorder[idx] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}
