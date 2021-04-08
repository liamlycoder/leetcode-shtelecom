package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyNode := &ListNode{Val: -1, Next: head}
	cur := dummyNode
	for cur.Next != nil && cur.Next.Next != nil {
		node1 := cur.Next
		node2 := cur.Next.Next
		cur.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		cur = node1
	}
	return dummyNode.Next
}
