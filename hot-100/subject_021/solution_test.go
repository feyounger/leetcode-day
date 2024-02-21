package subject_021

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	lists := mergeTwoLists(list1, list2)
	fmt.Println(lists)
}
