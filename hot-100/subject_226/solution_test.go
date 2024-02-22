package subject_226

import (
	"fmt"
	"testing"
)

func TestInvertTree(t *testing.T) {
	t1 := TreeNode{Val: 0, Left: &TreeNode{
		Val: 1,
	}, Right: &TreeNode{
		Val: 2,
	}}
	tree := invertTree(&t1)
	fmt.Println(tree)
}
