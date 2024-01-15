package subject_049

import (
	"sort"
	"strings"
)

/**
49.字母异位词分组
*/
func groupAnagrams(strs []string) [][]string {
	var result [][]string
	for i := 0; i < len(strs); i++ {
		if strings.HasPrefix(strs[i], "#") {
			continue
		}
		var temArr []string
		temArr = append(temArr, strs[i])
		front := quickSort(strs[i])
		for j := i + 1; j < len(strs); j++ {
			if len(strs[i]) != len(strs[j]) {
				continue
			}
			if front == quickSort(strs[j]) {
				temArr = append(temArr, strs[j])
				strs[j] = "#"
			}
		}
		result = append(result, temArr)
	}
	return result
}

func quickSort(str string) string {
	var arr []string
	kind := []byte(str)
	for _, s := range kind {
		arr = append(arr, string(s))
	}
	sort.Strings(arr)
	return strings.Join(arr, "")
}

func groupAnagrams1(strs []string) [][]string {
	var result [][]string
	for i := 0; i < len(strs); i++ {
		if strings.HasPrefix(strs[i], "#") {
			continue
		}
		var temArr []string
		temArr = append(temArr, strs[i])
		for j := i + 1; j < len(strs); j++ {
			temp := make(map[byte]byte)
			if len(strs[i]) != len(strs[j]) {
				continue
			}
			for _, by := range []byte(strs[i]) {
				temp[by] = by
			}
			frontLen := len(temp)
			for _, by := range []byte(strs[j]) {
				temp[by] = by
			}
			if len(temp) == frontLen {
				temArr = append(temArr, strs[j])
				strs[j] = "#"
			}
		}
		result = append(result, temArr)
	}
	return result
}

func groupAnagrams2(strs []string) [][]string {
	var result [][]string
	arrMap := map[string][]string{}
	for i := range strs {
		s := []byte(strs[i])
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		arrMap[string(s)] = append(arrMap[string(s)], strs[i])
	}
	for _, value := range arrMap {
		result = append(result, value)
	}
	return result
}
