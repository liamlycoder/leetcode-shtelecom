public class Solution_Sword_10_2 {
    public int numWays(int n) {
        int a = 1, b = 1;
        int c = 0;
        if (n <= 1) {
            return 1;
        }
        for (int i = 0; i < n-1; ++i) {
            c = (a + b) % 1000000007;
            a = b;
            b = c;
        }
        return c % 1000000007;
    }
}
