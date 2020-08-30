class Solution:
    def shortestPalindrome(self, s):
        reverse = s[::-1]
        for i in range(len(s) + 1):
            if s.startswith(reverse[i:]):
                return reverse[:i] + s


if __name__ == '__main__':
    result = Solution().shortestPalindrome("aacecaaa")
    print(result)
