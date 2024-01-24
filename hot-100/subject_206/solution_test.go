package subject_206

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	l1 := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	list := reverseList1(l1)
	fmt.Println(list)
}
