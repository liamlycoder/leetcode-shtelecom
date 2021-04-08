package main

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummyNode := &ListNode{Val: -1, Next: head}
	cur := dummyNode
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			deleteVal := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == deleteVal {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummyNode.Next
}
