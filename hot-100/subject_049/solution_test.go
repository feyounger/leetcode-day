package subject_049

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	str := "bac"
	var arr []string
	kind := []byte(str)
	for _, s := range kind {
		arr = append(arr, string(s))
	}
	sort.Strings(arr)
	fmt.Println(strings.Join(arr, ""))
}

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	t.Log(groupAnagrams2(strs))
}
