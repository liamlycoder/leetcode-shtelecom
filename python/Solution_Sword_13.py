# !/usr/bin/python
# coding: utf8
# Time: 2020/10/12 13:37
# Author: Liam
# E-mail: luyu.real@qq.com
# Software: PyCharm
class Solution:
    def movingCount(self, m: int, n: int, k: int) -> int:
        from queue import Queue
        q = Queue()
        result = set()
        q.put((0, 0))
        while not q.empty():
            x, y = q.get()
            if (x, y) not in result and 0 <= x < m and 0 <= y < n and Solution.check(x, y, k):
                result.add((x, y))
                for nx, ny in [(x+1, y), (x, y+1)]:
                    q.put((nx, ny))
        return len(result)

    @staticmethod
    def check(x, y, k):
        x = str(x)
        y = str(y)
        result = 0
        for num in x + y:
            result += int(num)
        return result <= k

if __name__ == '__main__':
    pass






