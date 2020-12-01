import entity.ListNode;

/**
 * @author liamcoder
 * @date 2020/12/1 14:11
 */
public class Solution141 {
    public boolean hasCycle(ListNode head) {
        if (head == null) {
            return false;
        }
        ListNode l1 = head;
        ListNode l2 = head.next;
        while (l1 != null && l2 != null && l2.next != null) {
            if (l1 == l2) {
                return true;
            } else {
                l1 = l1.next;
                l2 = l2.next.next;
            }
        }
        return false;
    }
}
