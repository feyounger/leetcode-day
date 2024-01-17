package subject_011

/*
*
11.盛最多水的容器
*/
func maxArea(height []int) int {
	var res int
	j := len(height) - 1
	i := 0
	for i < j {
		var hight int
		width := j - i
		if height[j] > height[i] {
			hight = height[i]
			i++
		} else {
			hight = height[j]
			j--
		}
		area := width * hight
		if area > res {
			res = area
		}
	}
	return res
}
