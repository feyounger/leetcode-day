package subject_024

import "testing"

func TestSwapPairs(t *testing.T) {
	l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	swapPairs(l)
}
