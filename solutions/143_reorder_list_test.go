package solutions

import (
	"reflect"
	"testing"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 2, 3, 4}, []int{1, 4, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}},
	}

	for _, test := range tests {
		head := SliceToList(test.input)
		ReorderList(head)
		result := ListToSlice(head)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("got %v, want %v", result, test.expected)
		}
	}
}

func SliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, num := range nums {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	return dummy.Next
}

func ListToSlice(head *ListNode) []int {
	var nums []int
	for node := head; node != nil; node = node.Next {
		nums = append(nums, node.Val)
	}
	return nums
}
