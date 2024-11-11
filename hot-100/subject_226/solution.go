package subject_226

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	leftNode := invertTree(root.Left)
	rightNode := invertTree(root.Right)
	root.Right = leftNode
	root.Left = rightNode
	return root

}
