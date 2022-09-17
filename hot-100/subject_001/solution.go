package subject_001

/**
1. 两数之和
*/
func twoSum01(array []int, target int) []int {
	var result []int
	for i, v1 := range array {
		for j := i + 1; j < len(array); j++ {
			if v1+array[j] == target {
				result = append(result, i)
				result = append(result, j)
				break
			}
		}
	}
	return result
}

func twoSum02(array []int, target int) []int {
	var result []int
	resultMap := make(map[int]int)
	for i, v1 := range array {
		index, ok := resultMap[v1]
		if ok {
			result = append(result, i, index)
			break
		}
		resultMap[target-v1] = i
	}
	return result
}

// execInstance 执行测试用例
func execInstance(array []int, target int) []int {
	return twoSum02(array, target)
}
