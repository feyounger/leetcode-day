package subject_876

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	l1 := head
	l2 := head
	for l2 != nil && l2.Next != nil {
		l1 = l1.Next
		l2 = l2.Next.Next

	}
	return l1
}
