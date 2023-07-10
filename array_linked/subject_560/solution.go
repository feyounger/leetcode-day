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
