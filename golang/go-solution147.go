package main

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummyHead := &ListNode{Val: -1, Next: head}
	lastSorted, cur := head, head.Next
	for cur != nil {
		if lastSorted.Val <= cur.Val {
			lastSorted = lastSorted.Next
		} else {
			find := dummyHead  // 该节点用于遍历寻找插入点
			for find.Next.Val <= cur.Val {
				find = find.Next
			}

			lastSorted.Next = cur.Next  // 把cur后面的节点接起来
			cur.Next = find.Next  // 下面两行就是将cur节点插入到find后面的基本操作
			find.Next = cur
		}
		cur = lastSorted.Next
	}
	return dummyHead.Next
}

