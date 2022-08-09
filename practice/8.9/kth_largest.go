/**
 * @Author: lihaoran
 * @Date: 2022/8/9 22:27
 * @Description:
 */

package __9

import "github.com/bill_lhr/personal/common"

func kthLargest(root *common.TreeNode, k int) int {
	stack := make([]*common.TreeNode, 0)
	cur := root
	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Right
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return node.Val
		}
		cur = node.Left
	}
	return 0
}
