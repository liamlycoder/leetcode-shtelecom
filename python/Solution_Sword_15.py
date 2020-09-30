class Solution:
    def hammingWeight(self, n):
        res = 0
        while n:
            res += 1
            n &= n - 1
        return res
