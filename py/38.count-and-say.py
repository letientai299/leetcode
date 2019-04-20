#
# @lc app=leetcode id=38 lang=python3
#
# [38] Count and Say
#
# https://leetcode.com/problems/count-and-say/description/
#
# algorithms
# Easy (40.06%)
# Total Accepted:    272.3K
# Total Submissions: 679.6K
# Testcase Example:  '1'
#
# The count-and-say sequence is the sequence of integers with the first five
# terms as following:
# 
# 
# 1.     1
# 2.     11
# 3.     21
# 4.     1211
# 5.     111221
# 
# 
# 1 is read off as "one 1" or 11.
# 11 is read off as "two 1s" or 21.
# 21 is read off as "one 2, then one 1" or 1211.
# 
# Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the
# count-and-say sequence.
# 
# Note: Each term of the sequence of integers will be represented as a
# string.
# 
# 
# 
# Example 1:
# 
# 
# Input: 1
# Output: "1"
# 
# 
# Example 2:
# 
# 
# Input: 4
# Output: "1211"
# 
#


class Solution:
    def countAndSay(self, n: int) -> str:
        s = '1'
        for i in range(1, n):
            cur = s[0]
            count = 1
            next_str = ''
            for c in s[1:]:
                if c == cur:
                    count += 1
                    continue

                next_str += str(count) + cur
                cur = c
                count = 1

            next_str += str(count) + cur
            s = next_str

        return s


if __name__ == '__main__':
    tests = [
        (1, '1'),
        (2, '11'),
        (3, '21'),
        (4, '1211'),
        (5, '111221'),
        (6, '312211'),
        (7, '13112221'),
        (8, '1113213211'),
    ]
    for tc in tests:
        param = tc[0]
        actual = Solution().countAndSay(param)
        expected = tc[1]
        if actual != expected:
            print(f"({param}) == {actual} != {expected}")
