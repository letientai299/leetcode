// Array Wrapper (Easy)
// https://leetcode.com/problems/array-wrapper/
//
// Create a class `ArrayWrapper` that accepts an array of integers in its
// constructor. This class should have two features:
//
// - When two instances of this class are added together with the `+` operator,
// the resulting value is the sum of all the elements in both arrays.
// - When the `String()` function is called on the instance, it will return a
// comma separated string surrounded by brackets. For example, `[1,2,3]`.
//
// **Example 1:**
//
// ```
// Input: nums = [[1,2],[3,4]], operation = "Add"
// Output: 10
// Explanation:
// const obj1 = new ArrayWrapper([1,2]);
// const obj2 = new ArrayWrapper([3,4]);
// obj1 + obj2; // 10
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [[23,98,42,70]], operation = "String"
// Output: "[23,98,42,70]"
// Explanation:
// const obj = new ArrayWrapper([23,98,42,70]);
// String(obj); // "[23,98,42,70]"
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [[],[]], operation = "Add"
// Output: 0
// Explanation:
// const obj1 = new ArrayWrapper([]);
// const obj2 = new ArrayWrapper([]);
// obj1 + obj2; // 0
// ```
//
// **Constraints:**
//
// - `0 <= nums.length <= 1000`
// - `0 <= nums[i] <= 1000`
// - `Note: nums is the array passed to the constructor`
//
// Langs: javascript, typescript
export class ArrayWrapper {
  private readonly nums: number[];
  constructor(nums: number[]) {
    this.nums = nums ?? [];
  }

  // in JS/TS, when we use the operator "+" on an unsupported value,
  // the engine will try to convert it to something that is supported.
  // hence, for this problem, we need to return the sum of the numbers.
  valueOf(): number {
    return this.nums.reduce((a, b) => a + b, 0);
  }

  toString(): string {
    return `[${this.nums.toString()}]`;
  }

  // this is slower than the built-in implementation
  _toString(): string {
    if (this.nums.length === 0) {
      return '[]';
    }

    let s = '[';
    s += this.nums[0]!.toString();
    for (let i = 1; i < this.nums.length; i++) {
      s += ',';
      s += this.nums[i]!.toString();
    }
    return s + ']';
  }
}

/**
 * const obj1 = new ArrayWrapper([1,2]);
 * const obj2 = new ArrayWrapper([3,4]);
 * obj1 + obj2; // 10
 * String(obj1); // "[1,2]"
 * String(obj2); // "[3,4]"
 */
