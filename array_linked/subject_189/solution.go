package subject_189

func rotate(nums []int, k int) {
	arr := make([]int, len(nums))
	copy(arr, nums)
	k = k % len(arr)
	for i := 0; i < len(arr); i++ {
		tmp := (len(arr) - k + i) % len(arr)
		nums[i] = arr[tmp]
	}
}

func rotate1(nums []int, k int) {
	k = k % len(nums)
	left := nums[len(nums)-k:]
	right := nums[:len(nums)-k]
	left = append(left, right...)
	nums = left
}
