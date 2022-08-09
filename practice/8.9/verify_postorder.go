/**
 * @Author: lihaoran
 * @Date: 2022/8/9 22:13
 * @Description:
 */

package __9

// 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。
func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	length := len(postorder)
	root := postorder[length-1]
	idx := 0
	for ; idx < length-1; idx++ {
		if postorder[idx] > root {
			break
		}
	}
	if idx < length-1 {
		for i := idx; i < length-1; i++ {
			if postorder[i] < root {
				return false
			}
		}
	}
	return verifyPostorder(postorder[:idx]) && verifyPostorder(postorder[idx:length-1])
}
