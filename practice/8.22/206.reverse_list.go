package __22

import "github.com/bill_lhr/personal/common"

func reverseList(head *common.ListNode) *common.ListNode {
	var pre *common.ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
