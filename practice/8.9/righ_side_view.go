/**
 * @Author: lihaoran
 * @Date: 2022/8/9 23:02
 * @Description:
 */

package __9

import "github.com/bill_lhr/personal/common"

func rightSideView(root *common.TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	queue := []*common.TreeNode{root}
	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[i]
			if i == length-1 {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[:length]
	}
	return res
}
