package solutions

func Insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)
	i := 0

	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval = merge(newInterval, intervals[i])
		i++
	}
	result = append(result, newInterval)

	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}

	return result
}

func merge(a, b []int) []int {
	return []int{min(a[0], b[0]), max(a[1], b[1])}
}
