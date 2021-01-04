/**
 * @author liamcoder
 * @date 2021/1/4 8:46
 */
public class Solution509 {
    public int fib(int n) {
        int x = 0, y = 1;
        if (n < 2) {
            return n;
        }
        int result = 0;
        for (int i = 1; i < n; ++i) {
            result = x + y;
            x = y;
            y = result;
        }
        return result;
    }
}
