package main

func partition2(head *ListNode, x int) *ListNode {
	small, large := &ListNode{}, &ListNode{}
	smallDummy, largeDummy := small, large
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}

	small.Next = largeDummy.Next
	large.Next = nil
	return smallDummy.Next
}
