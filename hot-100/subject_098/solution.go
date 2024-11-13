package subject_098

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return Re(root, math.MinInt64, math.MaxInt64)
}

func Re(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return Re(root.Left, lower, root.Val) && Re(root.Right, root.Val, upper)

}
