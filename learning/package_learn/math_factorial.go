package package_learn

func Math_Factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * Math_Factorial(n-1)
}
