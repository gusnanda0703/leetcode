package solutions

func LengthOfLongestSubstring(s string) int {
	mapRune := make(map[rune]int)
	start := 0
	maxCount := 0

	for i, v := range s {
		if lastI, ok := mapRune[v]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxCount {
			maxCount = i - start + 1
		}
		mapRune[v] = i
	}

	return maxCount
}
