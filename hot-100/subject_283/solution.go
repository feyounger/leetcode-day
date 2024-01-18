package subject_283

/*
*
283 移动零
*/
func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums) && j < len(nums); j++ {
		if nums[i] == 0 {
			if nums[j] != 0 {
				nums[i] = nums[j]
				nums[j] = 0
				i++
			} else {
				continue
			}
		} else {
			i++
		}
	}
}
