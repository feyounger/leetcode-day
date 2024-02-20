package subject_024

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	for head != nil && head.Next != nil {
		next := head.Next
		head.Next = next.Next
		next.Next = head
	}
	return head
}
