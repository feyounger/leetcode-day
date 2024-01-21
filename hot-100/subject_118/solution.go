package subject_118

func generate(numRows int) [][]int {
	res := [][]int{{1}}
	if numRows == 1 {
		return res
	}
	for i := 0; i < numRows-1; i++ {
		res = append(res, yh(res[i]))
	}
	return res
}
func yh(arr []int) []int {
	res := make([]int, len(arr)+1)
	res[0] = 1
	res[len(res)-1] = 1
	for i := 1; i < len(arr); i++ {
		res[i] = arr[i] + arr[i-1]
	}
	return res
}
