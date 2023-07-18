package subject_1094

func carPooling(trips [][]int, capacity int) bool {
	array := make([]int, 1002)
	for i := range trips {
		array[trips[i][1]+1] += trips[i][0]
		array[trips[i][2]+1] -= trips[i][0]
	}

	for i := 1; i < len(array); i++ {
		array[i] += array[i-1]
		if array[i] > capacity {
			return false
		}
	}
	return true
}
