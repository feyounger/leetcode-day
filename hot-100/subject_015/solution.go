package subject_015

import (
	"math"
	"sort"
)

/**
15. 三数之和
*/

func threeSum03(nums []int) [][]int {
	var result [][]int
	var j, k, tmp int
	tmp = -1
	sort.Ints(nums)
	j = len(nums) - 1
	k = j - 1
	for i := 0; i < j; {
		if k == i {
			i++
			k = j - 1
			continue
		}
		if nums[j] == tmp {
			j--
			k = j - 1
			continue
		}
		first := math.Abs(float64(nums[i]))
		if nums[j]+nums[k] == int(first) {
			result = append(result, []int{nums[i], nums[k], nums[j]})
			tmp = nums[j]
			j--
			k = j - 1
		} else if nums[j]+nums[k] > int(first) {
			k--
		} else {
			i++
		}
	}
	return result
}

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
	// 排序
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		// i大于0后直接结束
		if nums[i] > 0 {
			return result
		}
		// 给i去重,
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		//左右两个指针 若sum大于0移动right，否则移动left
		for left < right {
			num2 := nums[left]
			num3 := nums[right]
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 对left和right去重
				for left < right && nums[left] == num2 {
					left++
				}
				for left < right && nums[right] == num3 {
					right--
				}
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return result
}
func execInstance(nums []int) [][]int {
	return threeSum01(nums)
}
