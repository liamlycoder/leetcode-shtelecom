package main


// 归并两个有序链表
func mergeList(l1, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	cur := dummyNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return dummyNode.Next
}

func sortList(head *ListNode) *ListNode {
	// 链表为空或者只有一个元素的时候不用进行排序
	if head == nil || head.Next == nil {
		return head
	}

	// 下面用快慢指针来寻找需要分裂的中点位置
	slow, fast := head, head
	var preSlow *ListNode = nil  // preSlow保存了slow的前一个节点
	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 然后从中点分裂，左右两部分进行递归排序
	preSlow.Next = nil
	left := sortList(head)
	right := sortList(slow)
	return mergeList(left, right)
}


