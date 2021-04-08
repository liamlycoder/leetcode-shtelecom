package main


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Next: head}
	first, second := head, dummyNode
	// 让first先往前走n的距离，这样的话当first走到尽头的时候，second所指向的就是需要删除的那个节点
	for i := 0; i < n; i++ {
		first = first.Next
	}

	// 一起走，让second走到正确的位置
	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return dummyNode.Next
}

func main() {

}
