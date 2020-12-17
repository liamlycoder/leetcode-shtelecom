/**
 * @author liamcoder
 * @date 2020/12/17 14:22
 */
public class Solution121 {
    public int maxProfit(int[] prices) {
        int profit = 0;
        int min = Integer.MAX_VALUE;
        for (int price : prices) {
            if (price < min) {
                min = price;
            }
            profit = Math.max(profit, price - min);
        }
        return profit;
    }
}
