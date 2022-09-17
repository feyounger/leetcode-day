package subject_002

/**
2. 两数相加
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers01(l1 *ListNode, l2 *ListNode) *ListNode {
	var l4 *ListNode
	l3 := &ListNode{}
	l4 = l3
	var addNum, temp int
	for {
		if l1 == nil {
			l1 = &ListNode{}
		}
		if l2 == nil {
			l2 = &ListNode{}
		}
		addNum = l1.Val + l2.Val + temp
		if addNum >= 10 {
			addNum %= 10
			temp = 1
		} else {
			temp = 0
		}
		l3.Val = addNum
		if l1.Next == nil && l2.Next == nil {
			if temp == 1 {
				l3.Next = &ListNode{
					Val: 1,
				}
			}
			break
		}
		l1 = l1.Next
		l2 = l2.Next
		l3.Next = &ListNode{}
		l3 = l3.Next
	}
	return l4
}

// execInstance 执行测试用例
func execInstance(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers01(l1, l2)
}
