package subject_303

type NumArray2 struct {
	SumArray []int
}

func Constructor2(nums []int) NumArray {
	return NumArray{
		SumArray: nums,
	}
}

func (this *NumArray) SumRange2(left int, right int) int {
	var sum int
	for i := left; i < right+1; i++ {
		sum += this.SumArray[i]
	}
	return sum
}
