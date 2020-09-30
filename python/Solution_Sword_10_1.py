from functools import lru_cache


class Solution:
    @lru_cache(None)
    def fib(self, n):
        if n < 2:
            return n
        return (self.fib(n-1) + self.fib(n-2)) % 1000000007


if __name__ == '__main__':
    print(Solution().fib(1))