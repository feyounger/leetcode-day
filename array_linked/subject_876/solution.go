package subject_23

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	l1 := &ListNode{
		Next: head,
	}
	l2 := l1
	for l2 != nil {
		l1 = l1.Next
		if l2.Next == nil {
			l2 = l2.Next
		} else {
			l2 = l2.Next.Next
		}

	}
	return l1
}
