package solutions

func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// find the middle
	right, left := head, head
	for left != nil && left.Next != nil {
		left = left.Next.Next
		right = right.Next
	}

	// reverse
	var prev *ListNode
	for right != nil {
		next := right.Next
		right.Next = prev
		prev = right
		right = next
	}

	first, second := head, prev
	for second.Next != nil {
		tmp := first.Next
		first.Next = second
		first = tmp

		tmp = second.Next
		second.Next = first
		second = tmp
	}

}

// func ReorderList(head *ListNode) {
// 	list := []*ListNode{}

// 	for node := head; node != nil; node = node.Next {
// 		list = append(list, node)
// 		list = append(list, nil)
// 	}

// 	for i := 1; i < len(list)-i-1; i += 2 {
// 		list[i] = list[len(list)-i-1]
// 	}

// 	for i := 0; list[i] != nil; i++ {
// 		list[i].Next = list[i+1]
// 	}
// }
