package solutions

func FindMaxLength(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	counterMap := make(map[int]int, len(nums))
	maxLen, sum := 0, 0

	for idx, num := range nums {
		if num == 1 {
			sum += 1
		} else {
			sum -= 1
		}

		if sum == 0 {
			maxLen = max(maxLen, idx+1)
		} else if _, ok := counterMap[sum]; !ok {
			counterMap[sum] = idx
		} else {
			maxLen = max(maxLen, idx-counterMap[sum])
		}
	}

	return maxLen
}
