import entity.ListNode;

import java.util.ArrayList;

public class Solution_Sword_06 {
    public int[] reversePrint(ListNode head) {
        if (head == null) return new int[0];
        ArrayList<Integer> arrayList = new ArrayList<>();
        while (head != null) {
            arrayList.add(0, head.val);
            head = head.next;
        }
        int[] result = new int[arrayList.size()];
        for (int i = 0; i < arrayList.size(); ++i) {
            result[i] = arrayList.get(i);
        }
        return result;
    }
}
