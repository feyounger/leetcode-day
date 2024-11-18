package subject_105

import (
	"testing"
)

func Test_buildTree(t *testing.T) {
	tree := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	t.Log(tree)
}
