package solutions

import "sort"

func FindMinArrowShots(points [][]int) int {
	// Sort the intervals by end point
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	arrows := 1
	curr := points[0]

	for _, point := range points[1:] {
		// If the current interval does not overlap with the previous one, increment the arrow count
		if curr[1] < point[0] {
			arrows++
			curr = point
		}
	}

	return arrows
}
