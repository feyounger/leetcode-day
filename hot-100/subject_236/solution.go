package subject_236

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	m := make(map[int]*TreeNode)
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			m[root.Left.Val] = root
		}
		if root.Right != nil {
			m[root.Right.Val] = root
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	m2 := make(map[*TreeNode]int)
	slice_p := mapToSlice(m, p)
	slice_q := mapToSlice(m, q)
	for i := range slice_q {
		m2[slice_q[i]] = 0
	}
	for i, elem := range slice_p {
		if _, ok := m2[elem]; ok {
			return slice_p[i]
		}
	}

	return nil
}

func mapToSlice(m map[int]*TreeNode, s *TreeNode) []*TreeNode {
	arr := []*TreeNode{s}
	for {
		value, ok := m[s.Val]
		if ok {
			arr = append(arr, value)
			s = value
			continue
		} else {
			return arr
		}
	}
}
