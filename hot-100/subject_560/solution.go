package subject_560

func subarraySum(nums []int, k int) int {
	var res, pre int
	m := make(map[int]int)
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if value, ok := m[pre-k]; ok {
			res += value
		}
		m[pre]++
	}
	return res
}
