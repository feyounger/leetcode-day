package subject_438

import "testing"

func TestFindAnagrams(t *testing.T) {
	s := "cbaebabacd"
	p := "abc"
	t.Log(findAnagrams1(s, p))
}
