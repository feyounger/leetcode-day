package subject_025

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	group := reverseKGroup(l, 2)
	fmt.Println(group)
}
