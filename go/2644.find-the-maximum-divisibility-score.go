package main

// Find the Maximum Divisibility Score
//
// Easy
//
// https://leetcode.com/problems/find-the-maximum-divisibility-score
//
// You are given two integer arrays `nums` and `divisors`.
//
// The **divisibility score** of `divisors[i]` is the number of indices `j` such
// that `nums[j]` is divisible by `divisors[i]`.
//
// Return the integer `divisors[i]` with the **maximum** divisibility score. If
// multiple integers have the maximum score, return the smallest one.
//
// **Example 1:**
//
// **Input:** nums = [2,9,15,50], divisors = [5,3,7,2]
//
// **Output:** 2
//
// **Explanation:**
//
// The divisibility score of `divisors[0]` is 2 since `nums[2]` and `nums[3]`
// are divisible by 5.
//
// The divisibility score of `divisors[1]` is 2 since `nums[1]` and `nums[2]`
// are divisible by 3.
//
// The divisibility score of `divisors[2]` is 0 since none of the numbers in
// `nums` is divisible by 7.
//
// The divisibility score of `divisors[3]` is 2 since `nums[0]` and `nums[3]`
// are divisible by 2.
//
// As `divisors[0]`, `divisors[1]`, and `divisors[3]` have the same divisibility
// score, we return the smaller one which is `divisors[3]`.
//
// **Example 2:**
//
// **Input:** nums = [4,7,9,3,9], divisors = [5,2,3]
//
// **Output:** 3
//
// **Explanation:**
//
// The divisibility score of `divisors[0]` is 0 since none of numbers in `nums`
// is divisible by 5.
//
// The divisibility score of `divisors[1]` is 1 since only `nums[0]` is
// divisible by 2.
//
// The divisibility score of `divisors[2]` is 3 since `nums[2]`, `nums[3]` and
// `nums[4]` are divisible by 3.
//
// **Example 3:**
//
// **Input:** nums = [20,14,21,10], divisors = [10,16,20]
//
// **Output:** 10
//
// **Explanation:**
//
// The divisibility score of `divisors[0]` is 2 since `nums[0]` and `nums[3]`
// are divisible by 10.
//
// The divisibility score of `divisors[1]` is 0 since none of the numbers in
// `nums` is divisible by 16.
//
// The divisibility score of `divisors[2]` is 1 since `nums[0]` is divisible by
// 20.
//
// **Constraints:**
//
// - `1 <= nums.length, divisors.length <= 1000`
// - `1 <= nums[i], divisors[i] <= 109`
func maxDivScore(nums []int, divisors []int) int {
  best := -1
  score := 0
  for _, d := range divisors {
    curScore := 0
    for _, n := range nums {
      if n%d == 0 {
        curScore++
      }
    }

    if score > curScore {
      continue
    }

    if score < curScore ||
      (score == curScore && d < best) ||
      (best == -1) {
      score = curScore
      best = d
    }
  }
  return best
}
