package main


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	// 让first先往前走n的距离，这样的话当first走到尽头的时候，second所指向的就是需要删除的那个节点
	for i := 0; i < n; i++ {
		first = first.Next
	}

	for ; first != nil; first = first.Next {
		second = second.Next
	}

	second.Next = second.Next.Next

	return dummy.Next
}

func main() {

}
