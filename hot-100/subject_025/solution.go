package subject_025

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverseRang(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

func reverseRang(a, b *ListNode) *ListNode {
	var pre *ListNode
	cur := a
	for b != cur {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
