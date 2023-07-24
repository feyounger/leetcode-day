package subject_141

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	l1 := head
	l2 := head
	for l2 != nil && l2.Next != nil {
		l1 = l1.Next
		l2 = l2.Next.Next
		if l1 == l2 {
			return true
		}
	}
	return false
}
