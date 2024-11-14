package subject_230

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var arr []int
	Re(root, &arr)
	sort.Ints(arr)
	return arr[k-1]
}

func Re(root *TreeNode, arr *[]int) {
	if root != nil {
		*arr = append(*arr, root.Val)
	} else {
		return
	}
	Re(root.Right, arr)
	Re(root.Left, arr)
}
