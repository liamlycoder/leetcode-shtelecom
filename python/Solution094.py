# !/usr/bin/python
# coding: utf8
# Time: 2020/9/14 上午9:37
# Author: Liam
# E-mail: luyu.real@qq.com
# Software: PyCharm
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def inorderTraversal(self, root):
        if not root:
            return []
        return self.inorderTraversal(root.left) + [root.val] + self.inorderTraversal(root.right)


class Solution2:
    def inorderTraversal(self, root):
        result = []
        stack = []
        current = root
        while stack or current:
            while current:
                stack.append(current)
                current = current.left
            current = stack.pop()
            result.append(current.val)
            current = current.right
        return result
