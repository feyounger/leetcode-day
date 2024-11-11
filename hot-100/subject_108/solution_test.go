package subject_108

import (
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	bst := sortedArrayToBST([]int{1, 2, 3, 4, 5})
	t.Log(bst)
}
