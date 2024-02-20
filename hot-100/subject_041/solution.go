package subject_041

import "sort"

func firstMissingPositive(nums []int) int {
	var tmp int
	m := make(map[int]bool, len(nums))
	for _, value := range nums {
		m[value] = true
		if value > tmp {
			tmp = value
		}
	}
	for i := 1; i < tmp; i++ {
		value := m[i]
		if !value {
			return i
		}
	}
	return 1 + tmp
}

func firstMissingPositive1(nums []int) int {
	tmp := 1
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			continue
		}
		if nums[i] > tmp {
			return tmp
		} else if nums[i] == tmp {
			tmp++
		}
	}
	return tmp + 1
}
