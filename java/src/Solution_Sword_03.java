import java.util.HashMap;

public class Solution_Sword_03 {
    public int findRepeatNumber(int[] nums) {
        HashMap<Integer, Boolean> repeat = new HashMap<Integer, Boolean>();
        for (int num : nums) {
            if (repeat.get(num) == null) {
                repeat.put(num, true);
            } else {
                return num;
            }
        }
        return -1;
    }
}
