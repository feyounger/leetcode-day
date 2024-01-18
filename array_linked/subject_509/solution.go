package subject_509

func fib1(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}

func fib2(n int) int {
	var result int
	if n == 1 || n == 2 {
		return 1
	}
	dpInt1 := 0
	dpInt2 := 1
	for i := 2; i <= n; i++ {
		result = dpInt1 + dpInt2
		dpInt1 = dpInt2
		dpInt2 = result
	}
	return result
}
