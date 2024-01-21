package subject_1109

func corpFlightBookings(bookings [][]int, n int) []int {
	arr := make([]int, n+1)
	//res := make([]int, n+1)
	for i := range bookings {
		corpFlight(arr, bookings[i][0], bookings[i][1], bookings[i][2])
	}
	for i := 1; i < len(arr); i++ {
		arr[i] += arr[i-1]
	}
	return arr[1:]
}

func corpFlight(arr []int, left, right, num int) {
	arr[left] += num
	if right+1 < len(arr) {
		arr[right+1] -= num
	}
}
