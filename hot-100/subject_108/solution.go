package subject_108

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	node := recursion(&TreeNode{}, nums)
	return node
}

func recursion(tree *TreeNode, nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	middle := len(nums) / 2
	tree = &TreeNode{
		Val: nums[middle],
	}
	tree.Right = recursion(tree.Right, nums[middle+1:])
	tree.Left = recursion(tree.Left, nums[:middle])
	return tree
}
