package subject_438

import "sort"

// 暴力
func findAnagrams(s string, p string) []int {
	var result []int
	if len(p) > len(s) {
		return result
	}
	p = strSlice(p)
	k := len(p)
	for i := 0; i < len(s); i++ {
		if i+k > len(s) {
			break
		}
		pre := strSlice(s[i : i+k])
		if pre == p {
			result = append(result, i)
		}
	}
	return result
}

func strSlice(s string) string {
	tmp := []byte(s)
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})
	return string(tmp)
}

// 窗口移动
func findAnagrams1(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	// 如果s的长度小于p的长度，直接返回
	if sLen < pLen {
		return
	}
	//定义两个数组，分别记录s和p中各个字符出现的次数
	var sEng, pEng [26]int
	for i := range p {
		sEng[s[i]-'a']++
		pEng[p[i]-'a']++
	}
	// 如果s和p的前pLen个字符的出现次数相同，那么就是异位词
	if sEng == pEng {
		ans = append(ans, 0)
	}
	for i := 0; i < sLen-pLen; i++ {
		sEng[s[i]-'a']--
		sEng[s[i+pLen]-'a']++
		if sEng == pEng {
			ans = append(ans, i+1)
		}
	}

	return
}
