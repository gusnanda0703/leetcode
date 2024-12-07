package solutions

func UniquePaths(m int, n int) int {
	k := 1
	for i := 1; i <= m-1; i++ {
		k = k * (n - 1 + i) / i
	}
	return k
}
