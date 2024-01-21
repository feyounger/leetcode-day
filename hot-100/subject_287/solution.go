package subject_287

func findDuplicate(nums []int) int {
	m := make(map[int]int, len(nums))
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			return nums[i]
		}
		m[nums[i]] = 0
	}
	return 0
}
