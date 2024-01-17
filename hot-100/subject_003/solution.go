package subject_003

/*
*
3. 无重复字符的最长子串
*/
func lengthOfLongestSubstring01(s string) int {
	m := make(map[byte]int)
	var result, tmp int
	for i := 0; i < len(s); i++ {
		index, ok := m[s[i]]
		if ok {
			if tmp > result {
				result = tmp
			}
			tmp = 0
			i = index
			m = make(map[byte]int)
			continue
		}
		tmp++
		m[s[i]] = i
	}
	if tmp > result {
		result = tmp
	}
	return result
}

// execInstance 执行测试用例
func execInstance(s string) int {
	return lengthOfLongestSubstring01(s)
}
