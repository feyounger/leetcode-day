package subject_199

import (
	"testing"
)

func Test_rightSideView(t *testing.T) {
	node := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	view := rightSideView(&node)
	t.Log(view)
}
