package subject_086

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummy1 := &ListNode{Val: -1}
	dummy2 := &ListNode{Val: -1}
	p1 := dummy1
	p2 := dummy2
	p := head
	for p != nil {
		if p.Val >= x {
			p2.Next = &ListNode{Val: p.Val}
			p2 = p2.Next
		} else {
			p1.Next = &ListNode{Val: p.Val}
			p1 = p1.Next
		}
		p = p.Next
	}
	p1.Next = dummy2.Next
	return dummy1.Next
}
