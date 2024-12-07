package solutions

import (
	"reflect"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{2, 3}},
		{[]int{1, 1, 2}, []int{1}},
		{[]int{1}, []int{}},
	}

	for _, test := range tests {
		if got := FindDuplicates(test.nums); !reflect.DeepEqual(got, test.want) {
			t.Errorf("FindDuplicates(%v) = %v, want %v", test.nums, got, test.want)
		}
	}
}
