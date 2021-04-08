package main

func reorderList(head *ListNode)  {
	if head == nil {
		return
	}

	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList3(l2)
	mergeList2(l1, l2)
}

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList3(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func mergeList2(l1, l2 *ListNode)  {
	var tmp1, tmp2 *ListNode
	for l1 != nil && l2 != nil {
		tmp1 = l1.Next
		tmp2 = l2.Next

		l1.Next = l2
		l1 = tmp1

		l2.Next = l1
		l2 = tmp2
	}
}

func main() {

}
