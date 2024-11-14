package subject_199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var arr []int
	m := make(map[int]int)
	Re(root, m, 0)
	for i := 0; i < len(m); i++ {
		arr = append(arr, m[i])
	}
	return arr
}

func Re(root *TreeNode, arr map[int]int, i int) {
	if root != nil {
		arr[i] = root.Val
	} else {
		return
	}
	i++
	Re(root.Left, arr, i)
	Re(root.Right, arr, i)
}
