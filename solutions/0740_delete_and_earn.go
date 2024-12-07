package solutions

func DeleteAndEarn(nums []int) int {
	counts := make([]int, 10001)
	for _, num := range nums {
		counts[num] += num
	}

	prev, curr := 0, 0
	for _, count := range counts {
		prev, curr = curr, max(curr, prev+count)
	}

	return curr
}
