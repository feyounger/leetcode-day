package subject_101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return Symmetric(root.Left, root.Right)
}

func Symmetric(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return l.Val == r.Val && Symmetric(l.Left, r.Right) && Symmetric(l.Right, r.Left)
}
