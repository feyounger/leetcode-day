package subject_136

import "sort"

func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i-1 >= 0 && nums[i-1] == nums[i] {
			continue
		} else if i+1 < len(nums) && nums[i+1] == nums[i] {
			i++
			continue
		}
		return nums[i]
	}
	return 0
}

func singleNumber1(nums []int) int {
	var res int
	for _, num := range nums {
		res ^= num
	}
	return res
}
