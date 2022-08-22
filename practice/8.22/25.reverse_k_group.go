package __22

import "github.com/bill_lhr/personal/common"

func reverseKGroup(head *common.ListNode, k int) *common.ListNode {
	dummy := &common.ListNode{Next: head}
	left, right := dummy, dummy
	for {
		for i := 0; i < k && right != nil; i++ {
			right = right.Next
		}
		if right == nil {
			break
		}

		leftNext, rightNext := left.Next, right.Next
		right.Next = nil
		left.Next = reverse(leftNext)
		leftNext.Next = rightNext
		left, right = leftNext, leftNext
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
