package subject_23

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	target := findFromEnd(dummy, n+1)
	target.Next = target.Next.Next

	return dummy.Next

}

func findFromEnd(head *ListNode, n int) *ListNode {
	l1 := head
	l2 := head
	for i := 0; i < n; i++ {
		l1 = l1.Next
	}
	for l1 != nil {
		l1 = l1.Next
		l2 = l2.Next
	}
	return l2
}
