package invert_tree

import "github.com/bill_lhr/personal/common"

func invertTree(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	}
	queue := make([]*common.TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			node.Left, node.Right = node.Right, node.Left
		}
		queue = queue[0:length]
	}
	return root
}
