package main

// Distribute Money to Maximum Children
//
// Easy
//
// https://leetcode.com/problems/distribute-money-to-maximum-children
//
// You are given an integer `money` denoting the amount of money (in dollars)
// that you have and another integer `children` denoting the number of children
// that you must distribute the money to.
//
// You have to distribute the money according to the following rules:
//
// - All money must be distributed.
// - Everyone must receive at least `1` dollar.
// - Nobody receives `4` dollars.
//
// Return _the **maximum** number of children who may receive **exactly**_ `8`
// _dollars if you distribute the money according to the aforementioned rules_.
// If there is no way to distribute the money, return `-1`.
//
// **Example 1:**
//
// ```
// Input: money = 20, children = 3
// Output: 1
// Explanation:
// The maximum number of children with 8 dollars will be 1. One of the ways to
// distribute the money is:
// - 8 dollars to the first child.
// - 9 dollars to the second child.
// - 3 dollars to the third child.
// It can be proven that no distribution exists such that number of children
// getting 8 dollars is greater than 1.
//
// ```
//
// **Example 2:**
//
// ```
// Input: money = 16, children = 2
// Output: 2
// Explanation: Each child can be given 8 dollars.
//
// ```
//
// **Constraints:**
//
// - `1 <= money <= 200`
// - `2 <= children <= 30`
func distMoney(m int, c int) int {
  if m < c {
    return -1 // not enough money
  }

  if m == c {
    return 0 // all children get 1 dollar
  }

  // m>c

  if c == 0 {
    return -1 // still have money left
  }

  if c == 1 {
    if m == 4 {
      return -1 // invalid distribution
    }

    if m == 8 {
      return 1
    }

    return 0
  }

  if m < 8 {
    return 0
  }

  if r := distMoney(m-8, c-1); r != -1 {
    return 1 + r
  }

  best := -1
  for i := 1; i <= 7 && i != 4; i++ {
    if r := distMoney(m-i, c-1); r != -1 && best < r {
      best = r
    }
  }

  return best
}

// TODO (tai): O(7^n), can be much faster
