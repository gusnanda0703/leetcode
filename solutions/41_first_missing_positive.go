package solutions

func FirstMissingPositive(nums []int) int {
	i, n := 0, len(nums)
	for i < n {
		if 0 < nums[i] && nums[i] < n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}

	for i, v := range nums {
		if v != i+1 {
			return i + 1
		}
	}

	return n + 1
}
