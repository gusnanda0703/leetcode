package solutions

func IsPalindromeLinkedList(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// Find the middle of the list
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		// slow will +1 and fast is +2
		// so in the end of loop
		// slow alway 1/2 fast
		fast = fast.Next.Next
		slow = slow.Next
	}

	// Reverse the second half
	var prev *ListNode
	for slow != nil {
		// store next node for next loop
		next := slow.Next
		// point next to prev node
		slow.Next = prev
		// stor new node
		prev = slow
		// move to next node
		slow = next
	}

	// Compare the first and second half
	for prev != nil {
		if head.Val != prev.Val {
			return false
		}
		head = head.Next
		prev = prev.Next
	}

	return true
}
