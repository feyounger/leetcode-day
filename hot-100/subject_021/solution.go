package subject_021

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//定义一个哑节点
	var dummy = &ListNode{-1, nil}
	p1 := list1
	p2 := list2
	p := dummy
	//循环两个链表 直到某个链表为空
	for p1 != nil && p2 != nil {
		//判断两个链表中值的大小 谁小就入栈到结果集中，并前进
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		p = p.Next
	}
	// 若链表中有剩余直接追加到后面
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return dummy.Next
}
