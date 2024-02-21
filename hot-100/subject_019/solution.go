package subject_019

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	p1 := dummy
	p2 := dummy
	for i := 0; i < n+1; i++ {
		p1 = p1.Next
	}
	for p1 != nil {
		p2 = p2.Next
		p1 = p1.Next
	}
	p2.Next = p2.Next.Next
	return dummy.Next
}
