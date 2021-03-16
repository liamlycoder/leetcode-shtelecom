package main


// 翻转链表的代码，一定要背的滚瓜烂熟，（每天写一遍）
func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nextTemp := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextTemp
	}
	return pre
}

// 快慢指针找到链表中点
func findHalf(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}


func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	half := findHalf(head)
	reversedSecondHalf := reverseList(half.Next)

	p1 := head
	p2 := reversedSecondHalf
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return result
}

func main() {

}
