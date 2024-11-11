package subject_234

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	left := 0
	right := len(arr) - 1
	for left < right {
		if arr[left] != arr[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isPalindrome1(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	left, right := head, reverse(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cron := head
	for cron != nil {
		tmp := cron.Next
		cron.Next = pre
		pre = cron
		cron = tmp
	}
	return pre
}
