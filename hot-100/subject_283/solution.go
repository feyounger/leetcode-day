package subject_283

/*
*
283 移动零
*/
func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
}
