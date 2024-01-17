package subject_560

func subarraySum(nums []int, k int) int {
	var result int
	sumArray := make([]int, len(nums)+1)
	for i := 1; i < len(nums)+1; i++ {
		sumArray[i] = sumArray[i-1] + nums[i-1]
	}
	for i := 0; i < len(sumArray); i++ {
		for j := i + 1; j < len(sumArray); j++ {
			if sumArray[j]-sumArray[i] == k {
				result++
			}
		}
	}
	return result
}

// 思路： 用map的key记录前缀值，value记录出现的次数，用前缀值减去k，然后在hash中寻找，出现的次数即相加次数
func subarraySum1(nums []int, k int) int {
	var result, pre int
	m := make(map[int]int)
	// 特殊情况处理，虚拟增加差值为0的情况
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if count, ok := m[pre-k]; ok {
			result += count
		}
		m[pre]++
	}
	return result
}
