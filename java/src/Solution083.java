import entity.ListNode;


/**
 * @author liamcoder
 * @date 2020/12/30 15:43
 */
public class Solution083 {
    public ListNode deleteDuplicates(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }
        head.next = deleteDuplicates(head.next);
        return head.val == head.next.val ? head.next : head;
    }
}
