package solutions

import "sort"

func SortedSquares(nums []int) []int {
	s := make([]int, len(nums))
	for i, v := range nums {
		s[i] = v * v
	}
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

	return s
}
