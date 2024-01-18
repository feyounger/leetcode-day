package subject_005

import "reflect"

/**
最长回文子串
*/

// 中间扩散法
func longestPalindrome02(s string) string {
	if s == "" {
		return s
	}
	var start, end int
	for i := 0; i < len(s); i++ {
		// 单回文数
		left1, right1 := action(s, i, i)
		// 双回文素
		left2, right2 := action(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func action(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}

func longestPalindrome01(s string) string {
	var result, temp string
	var j int
	for i := 0; i < len(s); i++ {
		if j == len(s)-1 {
			break
		}
		temp = s[j : i+1]
		if validatePalindrome(temp) && len(temp) > len(result) {
			result = temp
		}
		if i == len(s)-1 {
			i = j
			j++
		}
	}
	if result == "" {
		result = string(s[0])
	}
	return result
}

func validatePalindrome(s string) bool {
	reverseStr := reverse(s)
	return reflect.DeepEqual(reverseStr, s)
}

func reverse(str string) string {
	length := len(str)
	if length == 0 {
		return ""
	}

	// 声明切片
	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, str[length-i-1])
	}
	return string(result)
}

func execInstance(s string) string {
	return longestPalindrome02(s)
}
