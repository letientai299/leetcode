namespace Src;

/*
 * @lc app=leetcode id=7 lang=csharp
 *
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (25.13%)
 * Total Accepted:    623.3K
 * Total Submissions: 2.5M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 *
 * Example 1:
 *
 *
 * Input: 123
 * Output: 321
 *
 *
 * Example 2:
 *
 *
 * Input: -123
 * Output: -321
 *
 *
 * Example 3:
 *
 *
 * Input: 120
 * Output: 21
 *
 *
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of
 * this problem, assume that your function returns 0 when the reversed integer
 * overflows.
 *
 */


public partial class Solution
{
  public int Reverse(int x)
  {
    var sign = 1;
    if (x < 0)
    {
      sign = -1;
      x = -x;

      if (x < 0) // still negative? must be overflowed
      {
        return 0;
      }
    }

    long v = 0;
    while (x != 0)
    {
      v = 10 * v + x % 10;
      if (v > int.MaxValue)
      {
        return 0;
      }

      x /= 10;
    }

    return (int)v * sign;
  }
}
