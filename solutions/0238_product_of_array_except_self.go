package solutions

func ProductExceptSelf(nums []int) []int {
	length := len(nums)
	output := make([]int, length)

	left := 1
	for i := 1; i < length; i++ {
		output[i] = left
		left *= nums[i-1]
	}

	right := 1
	for i := length - 1; i >= 0; i-- {
		output[i] *= right
		right *= nums[i]
	}

	return output
}
