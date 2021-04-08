package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{Next: head}
	pre := dummyHead
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummyHead.Next
			}
		}

		nextGroup := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = nextGroup
		pre = tail
		head = tail.Next
	}
	return dummyHead.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	var pre, cur *ListNode = nil, head
	for pre != tail {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return tail, head
}
