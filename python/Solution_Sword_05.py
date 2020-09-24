# !/usr/bin/python
# coding: utf8
# Time: 2020/9/24 上午10:04
# Author: Liam
# E-mail: luyu.real@qq.com
# Software: PyCharm
class Solution:
    def replaceSpace(self, s):
            return s.replace(" ", "%20")


if __name__ == '__main__':
    result = Solution().replaceSpace("I love you")
    print(result)

