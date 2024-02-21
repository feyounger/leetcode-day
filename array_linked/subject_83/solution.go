package subject_83

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	if slow != nil {
		slow.Next = nil
	}
	return head
}
