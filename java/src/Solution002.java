import entity.ListNode;

/**
 * @author liamcoder
 * @date 2021/1/4 9:24
 */
public class Solution002 {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode result = new ListNode(0);
        ListNode cur = result;
        int carry = 0;  // 表示进位
        while (l1 != null || l2 != null || carry != 0) {
            int l1val = l1 != null ? l1.val : 0;
            int l2val = l2 != null ? l2.val : 0;
            int sumVal = l1val +l2val + carry;
            carry = sumVal / 10;

            cur.next = new ListNode(sumVal % 10);
            cur = cur.next;

            if (l1 != null) {
                l1 = l1.next;
            }
            if (l2 != null) {
                l2 = l2.next;
            }
        }
        return result.next;
    }
}
