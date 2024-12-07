package solutions

func MaximumNumberOfIntegersToChooseFromARangeI(banned []int, n int, maxSum int) int {
	count := 0
	for i := 1; i < n+1; i++ {
		for _, b := range banned {
			if i == b {
				break
			}
		}
		maxSum -= i
		if maxSum > -1 {
			count++
		}
	}
	return count
}
