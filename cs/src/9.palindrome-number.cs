namespace Src;

/*
 * @lc app=leetcode id=9 lang=csharp
 *
 * [9] Palindrome Number
 *
 * https://leetcode.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (42.00%)
 * Total Accepted:    521.4K
 * Total Submissions: 1.2M
 * Testcase Example:  '121'
 *
 * Determine whether an integer is a palindrome. An integer is a palindrome
 * when it reads the same backward as forward.
 *
 * Example 1:
 *
 *
 * Input: 121
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it
 * becomes 121-. Therefore it is not a palindrome.
 *
 *
 * Example 3:
 *
 *
 * Input: 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a
 * palindrome.
 *
 *
 * Follow up:
 *
 * Coud you solve it without converting the integer to a string?
 *
 */

public partial class Solution
{
  public bool IsPalindrome(int x)
  {
    switch (x)
    {
      case < 0:
        return false; // single digit
      case < 10:
        return true;
    }

    var v = 0;
    var n = x;
    while (n > 0)
    {
      v = 10 * v + n % 10;
      n /= 10;
    }

    return v == x;
  }
}
