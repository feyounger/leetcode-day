package subject_1109

import "fmt"

func corpFlightBookings(bookings [][]int, n int) []int {
	array := make([]int, n+1)
	var tmp int
	for _, booking := range bookings {
		for i := booking[0]; i < booking[1]+1; i++ {
			for j := i; j < len(array); j++ {
				if i == j {
					tmp = booking[2]
				} else {
					tmp = 0
				}
				array[j] = array[j] + tmp
			}
		}
	}
	fmt.Println(array)
	for i := 1; i < len(array); i++ {
		array[i-1] = array[i] - array[i-1]
	}
	return array
}
