package __25

import "github.com/bill_lhr/personal/common"

func reverseKGroup(head *common.ListNode, k int) *common.ListNode {
	dummy := &common.ListNode{Next: head}
	slow, fast := dummy, dummy
	for {
		for i := 0; i < k && fast != nil; i++ {
			fast = fast.Next
		}
		if fast == nil {
			break
		}
		slowNext, fastNext := slow.Next, fast.Next
		fast.Next = nil
		slow.Next = reverse(slowNext)
		slowNext.Next = fastNext
		slow, fast = slowNext, slowNext
	}
	return dummy.Next
}

func reverse(node *common.ListNode) *common.ListNode {
	var pre *common.ListNode
	for node != nil {
		next := node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return pre
}
