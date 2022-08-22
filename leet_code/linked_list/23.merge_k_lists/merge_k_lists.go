package merge_k_lists

import "github.com/bill_lhr/personal/common"

/**
 * @Description:
	给你一个链表数组，每个链表都已经按升序排列。
	请你将所有链表合并到一个升序链表中，返回合并后的链表。
	链接：https://leetcode.cn/problems/merge-k-sorted-lists/
*/
func mergeKLists(lists []*common.ListNode) *common.ListNode {
	if len(lists) == 0 {
		return nil
	}
	var res *common.ListNode
	for _, list := range lists {
		res = merge(res, list)
	}
	return res
}

func merge(list1 *common.ListNode, list2 *common.ListNode) *common.ListNode {
	dummy := &common.ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return dummy.Next
}
