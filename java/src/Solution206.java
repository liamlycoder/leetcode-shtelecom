import entity.ListNode;

/**
 * @author liamcoder
 * @date 2020/12/30 9:59
 */
public class Solution206 {
    public ListNode reverseList(ListNode head) {
        ListNode pre = null;
        ListNode pCur = head;
        ListNode temp;
        while (pCur != null) {
            temp = pCur.next;
            pCur.next = pre;
            pre = pCur;
            pCur = temp;
        }
        return pre;
    }
}
