package subject_070

func climbStairs(n int) int {
	m := make(map[int]int)
	return climb(m, n)
}

func climb(m map[int]int, n int) int {
	if n == 1 || n == 2 {
		return n
	}
	if m[n] > 0 {
		return m[n]
	}
	m[n] = climb(m, n-1) + climb(m, n-2)
	return m[n]
}

func climbStairs1(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	var res int
	fi1 := 1
	fi2 := 2
	for i := 3; i <= n; i++ {
		res = fi1 + fi2
		fi1 = fi2
		fi2 = res
	}
	return res
}
