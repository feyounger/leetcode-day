package subject_543

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var maxlen int
	Tree(root, &maxlen)
	return maxlen
}

func Tree(root *TreeNode, maxlen *int) int {
	if root == nil {
		return 0
	}
	left := Tree(root.Left, maxlen)
	right := Tree(root.Right, maxlen)
	temp := left + right
	if temp > *maxlen {
		*maxlen = temp
	}
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
