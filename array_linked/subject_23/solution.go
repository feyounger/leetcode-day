package subject_23

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	dummy := &ListNode{Val: -1}
	p0 := dummy

	queue := make(PriorityQueue, 0)
	heap.Init(&queue)
	for i := range lists {
		if lists[i] != nil {
			heap.Push(&queue, lists[i])
		}
	}
	for queue.Len() > 0 {
		pop := heap.Pop(&queue).(*ListNode)
		p0.Next = pop
		if pop.Next != nil {
			heap.Push(&queue, pop.Next)
		}
		p0 = p0.Next
	}
	return dummy.Next
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	node := x.(*ListNode)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() any {
	node := *pq
	n := len(node)
	res := node[n-1]
	*pq = node[0 : n-1]
	return res

}
