package subject_238

func productExceptSelf(nums []int) []int {
	l1 := make([]int, len(nums))
	l2 := make([]int, len(nums))
	res := make([]int, len(nums))
	l1[0] = nums[0]
	l2[0] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		l1[i] = l1[i-1] * nums[i]
	}
	for i := 1; i < len(nums); i++ {
		l2[i] = l2[i-1] * nums[len(nums)-1-i]
	}
	res[0] = l2[len(nums)-2]
	res[len(nums)-1] = l1[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		res[i] = l1[i-1] * l2[len(nums)-i-2]
	}
	return res
}
