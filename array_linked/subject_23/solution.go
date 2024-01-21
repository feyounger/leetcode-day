package subject_23

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	result := &ListNode{}
	if len(lists) == 0 {
		return nil
	}

	return result
}
