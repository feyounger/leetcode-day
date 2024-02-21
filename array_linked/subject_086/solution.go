package subject_086

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	//dummy1存比x小的链
	dummy1 := &ListNode{Val: -1}
	dummy2 := &ListNode{Val: -1}
	p1 := dummy1
	p2 := dummy2
	p := head
	//循环链表
	for p != nil {
		if p.Val >= x {
			p2.Next = p
			p2 = p2.Next
		} else {
			p1.Next = p
			p1 = p1.Next
		}
		//不能直接 p=p.next 需要断开原链表中的节点next指针 因为p1&p2和p指向同样的链表，当p前进时，p1,p2指向会指向对方，最终形成环
		temp := p.Next
		p.Next = nil
		p = temp
	}
	//连接头尾
	p1.Next = dummy2.Next
	return dummy1.Next
}
