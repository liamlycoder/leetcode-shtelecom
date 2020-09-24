# !/usr/bin/python
# coding: utf8
# Time: 2020/9/24 上午9:26
# Author: Liam
# E-mail: luyu.real@qq.com
# Software: PyCharm
class Solution:
    def findRepeatNumber(self, nums):
        repeat = {}
        for num in nums:
            if num not in repeat:
                repeat[num] = 1
            else:
                return num


if __name__ == '__main__':
    result = Solution().findRepeatNumber([2, 3, 1, 0, 2, 5, 3])
    print(result)
