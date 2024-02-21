package subject_005

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res1 := palindrome(s, i, i)
		res2 := palindrome(s, i, i+1)
		if len(res1) > len(res) {
			res = res1
		}
		if len(res2) > len(res) {
			res = res2
		}
	}
	return res
}

func palindrome(s string, l int, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
