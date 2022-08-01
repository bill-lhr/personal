package flatten

import "github.com/bill_lhr/personal/common"

/**
 * @Description: 给你二叉树的根结点 root ，请你将它展开为一个单链表：
	展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
	展开后的单链表应该与二叉树 先序遍历 顺序相同。

	链接：https://leetcode.cn/problems/flatten-binary-tree-to-linked-list
*/

// 1. 递归
func flatten(root *common.TreeNode) {
	cur := &common.TreeNode{}
	var dfs func(node *common.TreeNode)
	dfs = func(node *common.TreeNode) {
		if node == nil {
			return
		}
		left, right := node.Left, node.Right
		cur.Left, cur.Right = nil, node
		cur = node
		dfs(left)
		dfs(right)
	}
	dfs(root)
}

// 2. 非递归
func flatten2(root *common.TreeNode) {
	if root == nil {
		return
	}
	stack := []*common.TreeNode{root}
	cur := &common.TreeNode{}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		cur.Left, cur.Right = nil, node
		cur = node
	}
}
