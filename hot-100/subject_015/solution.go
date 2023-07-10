package subject_015

import "sort"

/**
15. 三数之和
*/

func threeSum01(nums []int) [][]int {
	var result [][]int
	m := make(map[int]bool)
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if f1, f2 := m[nums[i]]; f1 && f2 {
			continue
		}
		ok := action(nums, i-1, i+1, nums[i], &result)
		if ok {
			m[nums[i]] = true
		} else {
			m[nums[i]] = false
		}
	}
	return result
}

func action(nums []int, left, right, mid int, result *[][]int) bool {
	var temp int
	var flag bool
	index := left
	for left >= 0 && right < len(nums) {
		temp = nums[left] + nums[right]
		if nums[left]+nums[right]+mid == 0 {
			*result = append(*result, []int{nums[left], nums[right], mid})
			flag = true
			right++
			left = index
			continue
		}
		if -temp > mid {
			right++
			left = index
			continue
		}
		if -temp < mid {
			left--
			continue
		}
	}
	return flag
}

func threeSum02(nums []int) [][]int {
	var result [][]int
	var temp int
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			temp = nums[i] + nums[j]
			for k := j + 1; k < len(nums); k++ {
				if -temp == nums[k] {
					result = append(result, []int{nums[i], nums[j], nums[k]})
					break
				}
			}
		}
	}
	return result
}
func execInstance(nums []int) [][]int {
	return threeSum01(nums)
}
