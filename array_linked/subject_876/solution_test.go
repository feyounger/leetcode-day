package subject_876

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	fmt.Println(middleNode(list1))
}
