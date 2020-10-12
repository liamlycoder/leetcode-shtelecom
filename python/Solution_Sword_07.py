# !/usr/bin/python
# coding: utf8
# Time: 2020/10/12 14:27
# Author: Liam
# E-mail: luyu.real@qq.com
# Software: PyCharm

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def buildTree(self, preorder: list, inorder: list) -> TreeNode:
        if len(preorder) != 0:
            ind = inorder.index(preorder[0])
            root = TreeNode(preorder[0])
            root.left = self.buildTree(preorder[1:ind+1], inorder[:ind])
            root.right = self.buildTree(preorder[ind+1:], inorder[ind+1:])
            return root


