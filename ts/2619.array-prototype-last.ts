// Array Prototype Last
//
// Easy
//
// https://leetcode.com/problems/array-prototype-last
//
// Write code that enhances all arrays such that you can call
// the `array.last()` method on any array and it will return the last element.
// If there are no elements in the array, it should return `-1`.
//
// You may assume the array is the output of `JSON.parse`.
//
// **Example 1:**
//
// ```
// Input: nums = [null, {}, 3]
// Output: 3
// Explanation: Calling nums.last() should return the last element: 3.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = []
// Output: -1
// Explanation: Because there are no elements, return -1.
//
// ```
//
// **Constraints:**
//
// - `arr` is a valid JSON array
// - `0 <= arr.length <= 1000`
declare global {
  interface Array<T> {
    last(): T | -1;
  }
}

Array.prototype.last = function () {
  return this.length ? this[this.length - 1] : -1;
};

/**
 * const arr = [1, 2, 3];
 * arr.last(); // 3
 */
