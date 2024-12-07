package solutions

func Fib(n int) int {
	if n <= 1 {
		return n
	}

	p, c := 1, 1
	for i := 2; i < n; i++ {
		p, c = c, p+c
	}

	return c
}
