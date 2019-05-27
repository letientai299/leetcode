#
# @lc app=leetcode id=101 lang=python3
#
# [101] Symmetric Tree
#
# https://leetcode.com/problems/symmetric-tree/description/
#
# algorithms
# Easy (43.28%)
# Total Accepted:    391K
# Total Submissions: 903.3K
# Testcase Example:  '[1,2,2,3,4,4,3]'
#
# Given a binary tree, check whether it is a mirror of itself (ie, symmetric
# around its center).
#
# For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
#
#
# ⁠   1
# ⁠  / \
# ⁠ 2   2
# ⁠/ \ / \
# 3  4 4  3
#
#
#
#
# But the following [1,2,2,null,3,null,3] is not:
#
#
# ⁠   1
# ⁠  / \
# ⁠ 2   2
# ⁠  \   \
# ⁠  3    3
#
#
#
#
# Note:
# Bonus points if you could solve it both recursively and iteratively.
#
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

from .tree_node import TreeNode


class Solution:
    def isSym(self, p: TreeNode, q: TreeNode) -> bool:
        if (p is None) != (q is None):
            return False

        if p is None:
            return True

        if p.val != q.val:
            return False

        return self.isSym(p.left, q.right) and self.isSym(p.right, q.left)

    def isSymmetric(self, root: TreeNode) -> bool:
        if root is None:
            return False

        return self.isSym(root.left, root.right)
