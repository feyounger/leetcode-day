package subject_102

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	node := TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	fmt.Println(levelOrder(&node))
}
