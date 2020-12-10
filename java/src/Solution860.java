/**
 * @author liamcoder
 * @date 2020/12/10 16:05
 */
public class Solution860 {
    public boolean lemonadeChange(int[] bills) {
        if (bills == null || bills.length == 0) {
            return true;
        }

        int five = 0, ten = 0;
        for (int bill : bills) {
            if (bill == 5) {
                ++five;
            } else if (bill == 10) {
                --five;
                ++ten;
            } else if (ten > 0) {
                --ten;
                --five;
            } else {
                five -= 3;
            }
            if (five < 0) {
                return false;
            }
        }
        return true;
    }
}
