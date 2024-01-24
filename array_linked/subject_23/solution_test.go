package subject_23

import (
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	l1 := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	l3 := &ListNode{2, &ListNode{6, &ListNode{7, nil}}}
	nodes := []*ListNode{l1, l2, l3}
	fmt.Println(mergeKLists(nodes))
}
