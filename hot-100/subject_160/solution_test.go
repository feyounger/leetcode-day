package subject_160

import "testing"

func TestName(t *testing.T) {
	l3 := &ListNode{7, &ListNode{8, &ListNode{9, nil}}}
	l1 := &ListNode{1, &ListNode{2, l3}}
	l2 := &ListNode{3, &ListNode{4, &ListNode{5, l3}}}

	getIntersectionNode(l1, l2)

}
