package subject_114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	nodes := iteration(root)
	for i := 1; i < len(nodes); i++ {
		pre, curr := nodes[i-1], nodes[i]
		pre.Left, pre.Right = nil, curr
	}
}

func iteration(root *TreeNode) []*TreeNode {
	var list []*TreeNode
	if root != nil {
		list = append(list, root)
		list = append(list, iteration(root.Left)...)
		list = append(list, iteration(root.Right)...)
	}
	return list
}
