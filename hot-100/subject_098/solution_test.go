package subject_098

import "testing"

func Test_isValidBST(t *testing.T) {
	node := TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	t.Log(isValidBST(&node))
}
