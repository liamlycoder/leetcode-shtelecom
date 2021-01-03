import entity.ListNode;

/**
 * @author liamcoder
 * @date 2021/1/3 3:24 下午
 */
public class Solution086 {
    public ListNode partition(ListNode head, int x) {
        ListNode lessHead = new ListNode(0);
        ListNode moreHead = new ListNode(0);

        if (head == null) {
            return null;
        }

        ListNode cur = head;
        ListNode curLess = lessHead;
        ListNode curMore = moreHead;

        while (cur != null) {
            if (cur.val < x) {
                curLess.next = new ListNode(cur.val);
                curLess = curLess.next;
            } else {
                curMore.next = new ListNode(cur.val);
                curMore = curMore.next;
            }
            cur = cur.next;
        }
        curLess.next = moreHead.next;
        return lessHead.next;
    }
}
