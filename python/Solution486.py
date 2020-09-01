class Solution:
    def PredictTheWinner(self, nums):
        """
        步骤：
        1. dp[i][j]表示当数组剩下的部分下标为i到j时，当前玩家与另一名玩家的分数之差的最大值，注意当前玩家不一定是先手
        2. 当i==j时，只有一个数字，当前玩家只能拿这个数字
        3. 当i<j时，当前玩家可以选择nums[i]或者nums[j]，然后轮到另一个玩家在数组剩下部分选择数字
        4. 在两种方案中，当前玩家会选择最优方案，即：
               dp[i][j] = max(nums[i] - dp[i+1][j], nums[j] - dp[i][j-1])
        """
        length = len(nums)
        dp = [[0]*length for _ in range(length)]
        for i, num in enumerate(nums):
            dp[i][i] = num
        for i in range(length - 2, -1, -1):
            for j in range(i+1, length):
                dp[i][j] = max(nums[i] - dp[i+1][j], nums[j] - dp[i][j-1])
        return dp[0][length - 1] >= 0


if __name__ == '__main__':
    result = Solution().PredictTheWinner([1, 5, 233, 7])
    print(result)

