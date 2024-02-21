package subject_146

type ListNode struct {
	Val  int
	Next *ListNode
}

type LRUCache struct {
	Capacity int
	M        map[int]int
	Link     *ListNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		M:        make(map[int]int),
		Link:     &ListNode{},
	}
}

func (this *LRUCache) Get(key int) int {
	value, ok := this.M[key]
	if !ok {
		return -1
	}
	return value
}

func (this *LRUCache) Put(key int, value int) {
	oldValue, ok := this.M[key]
	if ok {
		dummy := &ListNode{Next: this.Link}
		tmp := this.Link
		for tmp != nil {
			if oldValue == tmp.Val {
				oldHead := dummy.Next

			}
			tmp = tmp.Next
		}
	}
}
