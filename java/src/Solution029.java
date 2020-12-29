/**
 * @author liamcoder
 * @date 2020/12/29 9:11
 */
public class Solution029 {
    public int divide(int dividend, int divisor) {
        boolean sign = (dividend > 0) ^ (divisor > 0);
        int result = 0;
        if (dividend > 0) {
            dividend = -dividend;
        }
        if (divisor > 0) {
            divisor = -divisor;
        }

        while (dividend <= divisor) {
            int tempResult = -1;
            int tempDivisor = divisor;
            while (dividend <= (tempDivisor << 1)) {
                if (tempDivisor < (Integer.MIN_VALUE >> 1)) {
                    break;
                }
                tempResult = tempResult << 1;
                tempDivisor = tempDivisor << 1;
            }
            dividend = dividend - tempDivisor;
            result += tempResult;
        }

        if (!sign) {
            if (result <= Integer.MIN_VALUE) {
                return Integer.MAX_VALUE;
            }
            result = -result;
        }
        return result;
    }
}
