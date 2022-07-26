/**
 * @Author: lihaoran
 * @Date: 2022/7/26 22:05
 * @Description:
 */

package reverse_list

import "github.com/bill_lhr/personal/common"

func reverseList(head *common.ListNode) *common.ListNode {
	cur := head
	var pre *common.ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
