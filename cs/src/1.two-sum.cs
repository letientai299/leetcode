/*
 * @lc app=leetcode id=1 lang=csharp
 *
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (42.01%)
 * Total Accepted:    1.5M
 * Total Submissions: 3.6M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers, return indices of the two numbers such that they
 * add up to a specific target.
 * 
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 * 
 * Example:
 * 
 * 
 * Given nums = [2, 7, 11, 15], target = 9,
 * 
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 * 
 * 
 * 
 * 
 */

using System;
using System.Collections.Generic;
using System.IO;

public partial class Solution
{
    public int[] TwoSum(int[] nums, int k)
    {
        var m = new Dictionary<int, int>();
        for (var i = 0; i < nums.Length; i++)
        {
            var x = nums[i];
            if (m.ContainsKey(k - x))
            {
                return new[] {i, m[k-x]};
            }

            m[x] = i;
        }

        return new int[] { };
    }
}