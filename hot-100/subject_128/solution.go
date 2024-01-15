package subject_128

import "sort"

/*
*
128 最长连续序列
*/
func longestConsecutive(nums []int) int {
	var res, front int
	tmp := make(map[int]int)
	var arr, arrRes []int
	for _, num := range nums {
		if _, ok := tmp[num]; ok {
			continue
		} else {
			tmp[num] = num
			arr = append(arr, num)
		}
	}
	sort.Ints(arr)
	for _, value := range arr {
		if front+1 == value {
			res++
		} else {
			arrRes = append(arrRes, res)
			res = 1
		}
		front = value
	}
	arrRes = append(arrRes, res)
	sort.Ints(arrRes)
	if len(arrRes) != 0 {
		return arrRes[len(arrRes)-1]
	} else {
		return 0
	}
}
