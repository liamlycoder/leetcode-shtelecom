/**
 * @author liamcoder
 * @date 2020/12/17 13:43
 */
public class Solution714 {
    // 动态规划
    // dp[i][0]代表的是第i天手里没有股票的最大利润，
    // dp[i][1]代表的是第i天手里有股票的最大利润。
    public int maxProfit(int[] prices, int fee) {
        int n = prices.length;
        int [][] dp = new int[n][2];
        dp[0][0] = 0;
        dp[0][1] = -prices[0];
        for (int i = 1; i < n; i++) {
            dp[i][0] = Math.max(dp[i-1][0], dp[i-1][1] + prices[i] - fee);
            dp[i][1] = Math.max(dp[i-1][1], dp[i-1][0] - prices[i]);
        }
        return dp[n-1][0];
    }

    public int maxProfix2(int[] prices, int fee) {
        // 压缩版动态规划，相比上面省了空间
        int n = prices.length;
        int dp1 = 0;  // 手里没有股票
        int dp2 = -prices[0];  // 手里有股票
        for (int i = 0; i < n; ++i) {
            dp1 = Math.max(dp1, dp2 + prices[i] - fee);
            dp2 = Math.max(dp2, dp1 - prices[i]);
        }
        return dp1;
    }

    public int maxProfix3(int[] prices, int fee) {
        // 贪心算法：min保存的是当前最小的股价
        int n = prices.length;
        int min = prices[0];
        int profit = 0;
        for (int i = 0; i < n; ++i) {
            if (prices[i] < min) {
                min = prices[i];
            } else if (prices[i] - fee > min) {
                profit += (prices[i] - min - fee);
                min = prices[i] - fee;
            }
        }
        return profit;
    }
}
