package main

func removeElements(head *ListNode, val int) *ListNode {
	dummyNode := &ListNode{Val: -1, Next: head}
	pre := dummyNode
	cur := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return dummyNode.Next
}
