#
# @lc app=leetcode id=20 lang=python3
#
# [20] Valid Parentheses
#
# https://leetcode.com/problems/valid-parentheses/description/
#
# algorithms
# Easy (35.90%)
# Total Accepted:    527.4K
# Total Submissions: 1.5M
# Testcase Example:  '"()"'
#
# Given a string containing just the characters '(', ')', '{', '}', '[' and
# ']', determine if the input string is valid.
# 
# An input string is valid if:
# 
# 
# Open brackets must be closed by the same type of brackets.
# Open brackets must be closed in the correct order.
# 
# 
# Note that an empty string is also considered valid.
# 
# Example 1:
# 
# 
# Input: "()"
# Output: true
# 
# 
# Example 2:
# 
# 
# Input: "()[]{}"
# Output: true
# 
# 
# Example 3:
# 
# 
# Input: "(]"
# Output: false
# 
# 
# Example 4:
# 
# 
# Input: "([)]"
# Output: false
# 
# 
# Example 5:
# 
# 
# Input: "{[]}"
# Output: true
#


class Solution:
    def isValid(self, s: str) -> bool:
        # check if s is empty or None
        if not s:
            return True  # because empty string can be considered as valid

        stack = []
        for c in s:
            if c == '(' or c == '[' or c == '{':
                stack.append(c)
                continue

            if len(stack) == 0:
                return False

            last = stack[-1]
            if c == ')' and last != '(':
                return False

            if c == ']' and last != '[':
                return False

            if c == '}' and last != '{':
                return False

            stack.pop()

        return len(stack) == 0


if __name__ == '__main__':
    tests = {
        '()': True,
        '([])': True,
        '(()': False,
        ')': False,
    }
    for param, expected in tests.items():
        actual = Solution().isValid(param)
        if actual != expected:
            print(f"isValid({param}) == {actual} != {expected}")
