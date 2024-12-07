package solutions

import (
	"sort"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	lenNums := len(nums)
	half := lenNums / 2
	if lenNums%2 == 0 {
		return float64(nums[half-1]+nums[half]) / 2
	} else {
		return float64(nums[half])
	}
}
