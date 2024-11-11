package subject_102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		next := []*TreeNode{}
		res = append(res, []int{})
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		queue = next
	}
	return res
}
