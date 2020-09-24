# !/usr/bin/python
# coding: utf8
# Time: 2020/9/24 上午10:15
# Author: Liam
# E-mail: luyu.real@qq.com
# Software: PyCharm
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None
class Solution:
    def reversePrint(self, head):
        result = []
        while head:
            result.insert(0, head.val)
            head = head.next
        return result

