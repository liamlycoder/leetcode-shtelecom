class Solution:
    def numWays(self, n):
        # 其实也就是斐波拉契数列
        a, b = 1, 1
        sum = 0
        if n <= 1:
            return 1
        for i in range(n-1):
            sum = a + b
            a = b
            b = sum
        return sum % 1000000007


if __name__ == '__main__':
    # 1 1 2 3
    print(Solution().numWays(1))
