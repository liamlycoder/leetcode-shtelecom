package main

type ListNode struct {
    Val int
    Next *ListNode
}
 
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	var num1, num2 int
	flag := 0
	for l1 != nil || l2 != nil {
		num1, num2 = 0, 0
		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}
		sum := num1 + num2 + flag
		sum, flag = sum % 10, sum / 10

		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}

    }
    if flag > 0 {
    	tail.Next = &ListNode{Val: flag}
	}
	return head
}


func main() {

}
