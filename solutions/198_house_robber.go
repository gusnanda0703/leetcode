package solutions

func Rob(nums []int) int {

	prev, curr := 0, 0
	for _, v := range nums {
		prev, curr = curr, max(prev+v, curr)
	}

	return curr
}
