package solutions

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	l3 := head
	curry := 0
	for l1 != nil || l2 != nil {
		sum := curry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		curry = sum / 10
		l3.Next = &ListNode{sum % 10, nil}
		l3 = l3.Next
	}

	if curry > 0 {
		l3.Next = &ListNode{curry, nil}
	}

	return head.Next
}
