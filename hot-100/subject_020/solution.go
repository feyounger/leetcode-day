package subject_020

func isValid(s string) bool {
	var arr []string
	for i := 0; i < len(s); i++ {
		ca := Ca(string(s[i]))
		if ca != "" {
			if len(arr) == 0 || arr[len(arr)-1] != ca {
				return false
			} else {
				arr = arr[:len(arr)-1]
			}
		} else {
			arr = append(arr, string(s[i]))
		}
	}
	return len(arr) == 0
}

func Ca(str string) string {
	switch str {
	case ")":
		return "("
	case "}":
		return "{"
	case "]":
		return "["
	default:
		return ""
	}
}
