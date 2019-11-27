package main

func fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonaccibase(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	n1, n2, res := 1, 1, 0

	for i := 3; i <= n; i++ {
		res = n1 + n2
		n1, n2 = n2, res
	}
	return res
}
