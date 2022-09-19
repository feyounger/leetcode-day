package subject_004

import "sort"

/**
寻找两个正序数组的中位数
*/
func findMedianSortedArrays01(nums1 []int, nums2 []int) float64 {
	var nums []int
	var median int
	var result float64
	nums = append(nums1, nums2...)
	sort.Ints(nums)
	median = len(nums) % 2
	if median == 0 {
		result = (float64(nums[len(nums)/2]) + float64(nums[len(nums)/2-1])) / 2
	} else {
		result = float64(nums[len(nums)/2])
	}
	return result
}

// execInstance 执行测试用例
func execInstance(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays01(nums1, nums2)
}
