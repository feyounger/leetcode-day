package subject_437

import "testing"

func Test_pathSum1(t *testing.T) {
	node := TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   -2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:  2,
				Left: nil,
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val:  -3,
			Left: nil,
			Right: &TreeNode{
				Val:   11,
				Left:  nil,
				Right: nil,
			},
		},
	}
	t.Log(pathSum(&node, 8))
}

func Test_pathSum2(t *testing.T) {
	node := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   -2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   -3,
			Left:  nil,
			Right: nil,
		},
	}
	t.Log(pathSum(&node, -1))
}
