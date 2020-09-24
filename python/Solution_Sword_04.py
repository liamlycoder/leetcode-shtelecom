# !/usr/bin/python
# coding: utf8
# Time: 2020/9/24 上午9:51
# Author: Liam
# E-mail: luyu.real@qq.com
# Software: PyCharm
class Solution:
    def findNumberIn2DArray(self, matrix, target):
        if matrix is None or len(matrix) == 0:
            return False
        m, n = len(matrix), len(matrix[0])
        row, col = 0, n - 1
        while row < m and col >= 0:
            if target > matrix[row][col]:
                row += 1
            elif target < matrix[row][col]:
                col -= 1
            else:
                return True
        return False


if __name__ == '__main__':
    print(Solution().findNumberIn2DArray([
        [1, 4, 7, 11, 15],
        [2, 5, 8, 12, 19],
        [3, 6, 9, 16, 22],
        [10, 13, 14, 17, 24],
        [18, 21, 23, 26, 30]
    ], 20))
