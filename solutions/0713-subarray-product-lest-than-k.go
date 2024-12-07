package solutions

func NumSubarrayProductLessThanK(nums []int, k int) int {

	ans, left, prod := 0, 0, 1
	for right, v := range nums {
		prod *= v
		for prod >= k {
			prod /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}
