package subject_053

func maxSubArray(nums []int) int {
	var res int
	res = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}
