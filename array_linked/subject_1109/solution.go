package subject_1109

import "fmt"

func corpFlightBookings(bookings [][]int, n int) []int {
	array := make([]int, n+1)
	for _, booking := range bookings {
		array[booking[0]] += booking[2]
		if booking[1]+1 < n+1 {
			array[booking[1]+1] -= booking[2]
		}
	}
	fmt.Println(array)
	for i := 1; i < len(array); i++ {
		array[i] = array[i-1] + array[i]
	}
	return array[1:]
}
