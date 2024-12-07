package solutions

func FindDuplicates(nums []int) []int {
	res := make([]int, 0, 100)
	for _, v := range nums {
		if nums[abs(v)-1] < 0 {
			res = append(res, abs(v))
		}
		nums[abs(v)-1] *= -1
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
