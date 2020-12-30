import entity.ListNode;

/**
 * @author liamcoder
 * @date 2020/12/30 13:26
 */
public class Solution024 {
    public ListNode swapPairs(ListNode head) {
        if (head != null && head.next != null) {
            ListNode nextNode = head.next;
            head.next = swapPairs(nextNode.next);
            nextNode.next = head;
            return nextNode;
        }
        return head;
    }
}
