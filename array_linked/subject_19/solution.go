package subject_23

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	//找到 要删除节点的父节点(要对其进行操作)
	target := findFromEnd(dummy, n+1)
	//删除其节点
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
