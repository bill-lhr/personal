package __22

import "github.com/bill_lhr/personal/common"

func reverseBetween(head *common.ListNode, left int, right int) *common.ListNode {
	dummy := &common.ListNode{Next: head}
	l, r := dummy, dummy // l 反转区间开始节点的前一个节点 r 反转区间的结束节点
	for i := 0; i < left; i++ {
		l = l.Next
		r = r.Next
	}
	for i := left; i <= right; i++ {
		r = r.Next
	}

	lNext, rNext := l.Next, r.Next
	r.Next = nil
	l.Next = reverse(lNext)
	lNext.Next = rNext

	return dummy.Next
}
