package subject_142

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	//先找出快慢链表的相遇点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	//没有像雨点直接返回
	if fast == nil || fast.Next == nil {
		return nil
	}
	//重新定义一个k链表指向head，将k和fast以同样的速度运行相等时就是相遇点。
	//假设slow走了x步遇到fast，那fast就走了2x，多走的x就是在圈里。
	k := head
	for k != fast {
		k = k.Next
		fast = fast.Next
	}
	return k
}
