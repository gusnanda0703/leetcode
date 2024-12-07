package solutions

import "testing"

func TestMinimumSize(t *testing.T) {
	teste := []struct {
		name   string
		nums   []int
		maxOps int
	}{
		{
			name:   "test case 1",
			nums:   []int{9},
			maxOps: 2,
		},
		{
			name:   "test case 2",
			nums:   []int{2, 4, 8, 2},
			maxOps: 4,
		},
	}

	for _, tt := range teste {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumSize(tt.nums, tt.maxOps)
			t.Logf("got: %d", got)
		})
	}
}
