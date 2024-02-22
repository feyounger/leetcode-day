package subject_148

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	var arr []int
	sortHead := head
	dummy := &ListNode{}
	dummy.Next = head
	for sortHead != nil {
		arr = append(arr, sortHead.Val)
		sortHead = sortHead.Next
	}
	sort.Ints(arr)
	for _, value := range arr {
		head.Val = value
		head = head.Next
	}
	return dummy.Next
}
