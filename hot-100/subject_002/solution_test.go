package subject_002

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers01(t *testing.T) {
	l1 := ListNode{
		Val: 2,
	}
	l1.Next = &ListNode{
		Val: 4,
	}
	l1.Next.Next = &ListNode{
		Val: 3,
	}

	l2 := ListNode{
		Val: 5,
	}
	l2.Next = &ListNode{
		Val: 6,
	}
	l2.Next.Next = &ListNode{
		Val: 7,
	}

	l3 := execInstance(&l1, &l2)
	fmt.Println(l3)

}

func TestAddTwoNumbers02(t *testing.T) {
	l1 := ListNode{
		Val: 9,
	}
	l1.Next = &ListNode{
		Val: 9,
	}
	l1.Next.Next = &ListNode{
		Val: 9,
	}
	l1.Next.Next.Next = &ListNode{
		Val: 9,
	}

	l2 := ListNode{
		Val: 9,
	}
	l2.Next = &ListNode{
		Val: 9,
	}
	l2.Next.Next = &ListNode{
		Val: 9,
	}

	l3 := execInstance(&l1, &l2)
	fmt.Println(l3)

}

func TestAddTwoNumbers03(t *testing.T) {
	l1 := ListNode{
		Val: 0,
	}

	l2 := ListNode{
		Val: 0,
	}

	l3 := execInstance(&l1, &l2)
	fmt.Println(l3)

}
