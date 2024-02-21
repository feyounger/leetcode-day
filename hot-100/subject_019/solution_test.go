package subject_019

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	l1 := &ListNode{1, nil}
	end := removeNthFromEnd(l1, 1)
	fmt.Println(end)
}
