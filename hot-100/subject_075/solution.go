package subject_075

import "fmt"

func sortColors(nums []int) {
	m := make(map[int]int, 3)
	var count int
	for i := range nums {
		m[nums[i]]++
	}
	for i := 0; i < 3; i++ {
		if num, ok := m[i]; ok {
			for j := 0; j < num; j++ {
				nums[count] = i
				count++
			}
		}
	}
	fmt.Println(nums)
}
