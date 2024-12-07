package solutions

func MinCostClimbingStairs(cost []int) int {
	prev, curr := 0, 0
	for _, v := range cost {
		prev, curr = curr, min(prev, curr)+v
	}

	return min(prev, curr)
}
