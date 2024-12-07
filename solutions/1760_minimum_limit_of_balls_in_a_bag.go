package solutions

func minimumSize(nums []int, maxOperations int) int {
	left, right := 1, 0
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for _, num := range nums {
		right = max(right, num)
	}
	check := func(nums []int, cost int, maxOperations int) bool {
		count := 0
		for _, v := range nums {
			count += (v - 1) / cost
			if count > maxOperations {
				return false
			}
		}
		return true
	}
	for left <= right {
		mid := (left + right) / 2
		if check(nums, mid, maxOperations) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
