package subject_148

import "testing"

func TestSortList(t *testing.T) {
	l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	sortList(l)
}
