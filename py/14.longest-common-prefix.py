#
# @lc app=leetcode id=14 lang=python3
#
# [14] Longest Common Prefix
#
# https://leetcode.com/problems/longest-common-prefix/description/
#
# algorithms
# Easy (33.00%)
# Total Accepted:    417.5K
# Total Submissions: 1.3M
# Testcase Example:  '["flower","flow","flight"]'
#
# Write a function to find the longest common prefix string amongst an array of
# strings.
# 
# If there is no common prefix, return an empty string "".
# 
# Example 1:
# 
# 
# Input: ["flower","flow","flight"]
# Output: "fl"
# 
# 
# Example 2:
# 
# 
# Input: ["dog","racecar","car"]
# Output: ""
# Explanation: There is no common prefix among the input strings.
# 
# 
# Note:
# 
# All given inputs are in lowercase letters a-z.
# 
#
from typing import List


class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if not strs or len(strs) == 0:
            return ''

        if len(strs) == 1:
            return strs[0]

        max_possible_prefix_length = min([len(s) for s in strs])

        i = 0
        while i < max_possible_prefix_length:
            subset = set([s[i] for s in strs])
            if len(subset) != 1:
                return strs[0][:i]
            i += 1

        return strs[0][:max_possible_prefix_length]


if __name__ == '__main__':
    tests = [
        (None, ''),
        ([], ''),
        ([''], ''),
        (['aaa'], 'aaa'),
        (['a', 'aa'], 'a'),
        (['aa', 'a'], 'a'),
        (['flower', 'flow', 'flight', 'fly'], 'fl'),
        (['flower', 'flow', 'flight', ''], ''),
        (['flower', 'flow', 'floor'], 'flo'),
    ]

    for param, expected in tests:
        actual = Solution().longestCommonPrefix(param)
        if actual != expected:
            print(f"longestCommonPrefix({param}) == {actual} != {expected}")
