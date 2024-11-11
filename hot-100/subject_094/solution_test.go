package subject_094

import (
	"fmt"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	node := TreeNode{
		Val: 1,
		//Left:  &TreeNode{Val: 2},
		//Right: &TreeNode{Val: 3},
	}
	fmt.Println(inorderTraversal(&node))
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	center := arr[0]
	var left, right []int
	for _, value := range arr[1:] {
		if value <= center {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, center), right...)
}

func TestQuick(t *testing.T) {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	fmt.Println("Unsorted array:", arr)
	fmt.Println("Sorted array:", quickSort(arr))
}
