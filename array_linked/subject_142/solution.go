package subject_142

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	p1 := head
	p2 := head
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			break
		}
	}
	if p2 == nil || p2.Next == nil {
		return nil
	}
	p3 := head
	for p2 != p3 {
		p2 = p2.Next
		p3 = p3.Next
	}
	return p2
}
