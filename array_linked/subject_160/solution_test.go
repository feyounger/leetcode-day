package subject_160

import (
	"fmt"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	//common := &ListNode{
	//	Val: 8,
	//	Next: &ListNode{
	//		Val: 4,
	//		Next: &ListNode{
	//			Val: 5,
	//		},
	//	},
	//}
	list1 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	list2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	fmt.Println(getIntersectionNode(list1, list2))
}
