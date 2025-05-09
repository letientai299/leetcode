// Chunk Array (Easy)
// https://leetcode.com/problems/chunk-array
//
// Given an array `arr` and a chunk size `size`, return a **chunked** array.
// A **chunked** array contains the original elements in `arr`, but consists of
// subarrays each of length `size`. The length of the last subarray may be less
// than `size` if `arr.length` is not evenly divisible by `size`.
// You may assume the array is the output of `JSON.parse`. In other words, it is
// valid JSON.
// Please solve it without using lodash's `_.chunk` function.
// **Example 1:**
// ```
// Input: arr = [1,2,3,4,5], size = 1
// Output: [[1],[2],[3],[4],[5]]
// Explanation: The arr has been split into subarrays each with 1 element.
// ```
// **Example 2:**
// ```
// Input: arr = [1,9,6,3,2], size = 3
// Output: [[1,9,6],[3,2]]
// Explanation: The arr has been split into subarrays with 3 elements. However,
// only two elements are left for the 2nd subarray.
// ```
// **Example 3:**
// ```
// Input: arr = [8,5,3,2,6], size = 6
// Output: [[8,5,3,2,6]]
// Explanation: Size is greater than arr.length thus all elements are in the
// first subarray.
// ```
// **Example 4:**
// ```
// Input: arr = [], size = 1
// Output: []
// Explanation: There are no elements to be chunked so an empty array is
// returned.
// ```
// **Constraints:**
// - `arr` is a valid JSON array
// - `2 <= JSON.stringify(arr).length <= 105`
// - `1 <= size <= arr.length + 1`
// Langs: javascript, typescript
type JSONValue =
  | null
  | boolean
  | number
  | string
  | JSONValue[]
  | { [key: string]: JSONValue };
type Obj = Record<string, JSONValue> | Array<JSONValue>;

export function chunk<Obj>(arr: Obj[], size: number): Obj[][] {
  return Array.from({ length: Math.ceil(arr.length / size) }).map(
    (_, i) => arr.slice(i * size, (i + 1) * size),
    // this shorter method mutates the original array, so it's not desirable in practice
    // arr.splice(0, size),
  );
}
