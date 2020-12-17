/**
 * @author liamcoder
 * @date 2020/12/17 14:36
 */
public class Solution122 {
    public int maxProfit(int[] prices) {
        int n = prices.length;
        int profit = 0;
        int min = prices[0];
        for (int i = 1; i < n; ++i) {
            if (prices[i] < min) {
                min = prices[i];
            } else if (prices[i] > min) {
                profit += (prices[i] - min);
                min = prices[i];
            }
        }
        return profit;
    }
}
