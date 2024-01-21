package subject_303

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 * [[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
 */

// NumArray 303. 区域和检索 - 数组不可变
type NumArray struct {
	SumArray []int
}

func Constructor(nums []int) NumArray {
	sumArray := make([]int, len(nums)+1)
	for i := 1; i < len(sumArray); i++ {
		sumArray[i] = sumArray[i-1] + nums[i-1]
	}
	return NumArray{sumArray}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.SumArray[right+1] - this.SumArray[left]
}
