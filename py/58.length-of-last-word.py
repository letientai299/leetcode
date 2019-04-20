#
# @lc app=leetcode id=58 lang=python3
#
# [58] Length of Last Word
#
# https://leetcode.com/problems/length-of-last-word/description/
#
# algorithms
# Easy (32.15%)
# Total Accepted:    249.6K
# Total Submissions: 776.1K
# Testcase Example:  '"Hello World"'
#
# Given a string s consists of upper/lower-case alphabets and empty space
# characters ' ', return the length of last word in the string.
# 
# If the last word does not exist, return 0.
# 
# Note: A word is defined as a character sequence consists of non-space
# characters only.
# 
# Example:
# 
# Input: "Hello World"
# Output: 5
#


class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        i = len(s) - 1
        while i >= 0 and s[i] == ' ':
            i -= 1

        end = i
        while i >= 0 and s[i] != ' ':
            i -= 1

        return end - i


if __name__ == '__main__':
    tests = [
        ('   world   ', 5),
        ('world   ', 5),
        ('world', 5),
        ('Hello world', 5),
        ('', 0),
        ('    ', 0),
    ]
    for tc in tests:
        actual = Solution().lengthOfLastWord(tc[0])
        expected = tc[1]
        if actual != expected:
            print(f"lengthOfLastWord({tc[0]}) == {actual} != {expected}")
