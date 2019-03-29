#
# @lc app=leetcode id=374 lang=python
#
# [374] Guess Number Higher or Lower
#
# https://leetcode.com/problems/guess-number-higher-or-lower/description/
#
# algorithms
# Easy (38.93%)
# Total Accepted:    101.4K
# Total Submissions: 260.5K
# Testcase Example:  '10\n6'
#
# We are playing the Guess Game. The game is as follows:
#
# I pick a number from 1 to n. You have to guess which number I picked.
#
# Every time you guess wrong, I'll tell you whether the number is higher or
# lower.
#
# You call a pre-defined API guess(int num) which returns 3 possible results
# (-1, 1, or 0):
#
#
# -1 : My number is lower
# ⁠1 : My number is higher
# ⁠0 : Congrats! You got it!
#
#
# Example :
#
#
#
# Input: n = 10, pick = 6
# Output: 6
#
#
#
#
# The guess API is already defined for you.
# @param num, your guess
# @return -1 if my number is lower, 1 if my number is higher, otherwise return 0


class Solution(object):
    def guessNumber(self, n):
        """
        :type n: int
        :rtype: int
        """
        l = 1
        r = n
        while l <= r:
            m = l + (r-l)//2
            state = guess(m)
            if state == 0:
                return m

            if state == -1:
                r = m-1
            else:
                l = m+1

        return -1


#  x = 2
#  n = 2


#  def guess(num):
    #  if x < num:
        #  return -1
    #  elif x > num:
        #  return 1
    #  return 0


#  result = Solution().guessNumber(n)
#  print(result)
