package subject_21

//ListNode 21. 合并两个有序链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	l3 := dummy
	l1 := list1
	l2 := list2
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			l3.Next = l2
			l2 = l2.Next
		} else {
			l3.Next = l1
			l1 = l1.Next
		}
		l3 = l3.Next
	}
	if l1 != nil {
		l3.Next = l1
	}
	if l2 != nil {
		l3.Next = l2
	}
	return dummy.Next
}
