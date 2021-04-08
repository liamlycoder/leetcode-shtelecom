package main

func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode = nil
	var cur *ListNode = head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func main() {

}
