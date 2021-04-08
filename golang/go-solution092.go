package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Val: -1, Next: head}
	pre := dummyNode
	// 1. 把pre节点移动到待翻转的起始位置的前一个，即向后走left-1个位置
	for i := 0; i < left - 1; i++ {
		pre = pre.Next
	}

	// 2. 从pre所在位置走right - left + 1个位置，使得rightNode指向待翻转的右节点
	rightNode := pre
	for i := 0; i < right - left + 1; i++ {
		rightNode = rightNode.Next
	}

	// 3. 将待翻转的子链表截取出来
	leftNode := pre.Next
	cur := rightNode.Next
	pre.Next = nil
	rightNode.Next = nil

	// 4. 使用翻转链表算法进行翻转
	reverseLinkedList(leftNode)

	// 5. 将翻转后的子链表补回到原链表中去
	pre.Next = rightNode
	leftNode.Next = cur
	return dummyNode.Next
}

func reverseLinkedList(head *ListNode) {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}




func main() {

}
