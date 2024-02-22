package subject_104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var result, dep int
	var depthFunc func(*TreeNode)
	depthFunc = func(node *TreeNode) {
		if node == nil {
			return
		}
		dep++
		if node.Left == nil && node.Right == nil {
			if dep > result {
				result = dep
			}
		}
		depthFunc(node.Left)
		depthFunc(node.Right)
		dep--
	}
	depthFunc(root)
	return result
}

func maxDepth1(root *TreeNode) int {
	var depthFunc func(*TreeNode) int
	depthFunc = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDep := depthFunc(node.Left)
		rightDep := depthFunc(node.Right)
		return max(leftDep, rightDep) + 1
	}

	return depthFunc(root)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
