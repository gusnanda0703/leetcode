package solutions

func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	for i := range s {
		odd := expandPalindrome(s, i, i)
		even := expandPalindrome(s, i, i+1)
		maxLen := max(odd, even)
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func expandPalindrome(s string, l, r int) int {
	for 0 <= l && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}
