package subject_042

/*
*
42.接雨水
*/
// 思路：遍历height，找右边最大的值，若这个值比i的值大，结束寻找，计算i到这个值之间的雨水量(用找到的值和i的值中的最小值做基数),跳到这个值的位置继续遍历
func trap(height []int) int {
	var result int
	for i := 0; i < len(height); {
		if i+1 < len(height) && height[i] < height[i+1] {
			i++
			continue
		}
		var tmp, max int
		for j := i + 1; j < len(height); j++ {
			if height[j] >= height[i] {
				max = j
				tmp = height[i]
				break
			} else if height[j] >= tmp {
				max = j
				tmp = height[j]
			}
		}
		if height[i] < tmp {
			tmp = height[i]
		}
		k := i + 1
		for ; k < max; k++ {
			result += tmp - height[k]
		}
		i = k
	}
	return result
}

// 前后缀分解 当前i的雨水量是由左右两边最大值中的最小值减去当前值得到的
func trap1(height []int) int {
	var result int
	if len(height) == 1 {
		return 0
	}

	//获取前缀最大值
	pre := make([]int, len(height))
	pre[0] = height[0]
	for i := 1; i < len(height); i++ {
		pre[i] = max(pre[i-1], height[i])
	}
	//获取后缀最大值
	sub := make([]int, len(height))
	sub[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		sub[i] = max(sub[i+1], height[i])
	}
	//用前缀和后缀的最小值减去当前值，即为当前位置的雨水量
	for i := 0; i < len(height); i++ {
		result += min(pre[i], sub[i]) - height[i]
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
