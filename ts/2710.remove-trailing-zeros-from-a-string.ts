// Remove Trailing Zeros From a String (Easy)
// https://leetcode.com/problems/remove-trailing-zeros-from-a-string/
//
// Given a **positive** integer `num` represented as a string, return *the
// integer* `num` *without trailing zeros as a string*.
//
// **Example 1:**
//
// ```
// Input: num = "51230100"
// Output: "512301"
// Explanation: Integer "51230100" has 2 trailing zeros, we remove them and
// return integer "512301".
// ```
//
// **Example 2:**
//
// ```
// Input: num = "123"
// Output: "123"
// Explanation: Integer "123" has no trailing zeros, we return integer "123".
// ```
//
// **Constraints:**
//
// - `1 <= num.length <= 1000`
// - `num` consists of only digits.
// - `num` doesn't have any leading zeros.
//
// Langs: cpp, java, python, python3, c, csharp, javascript, typescript, php,
// swift, kotlin, dart, golang, ruby, scala, rust, racket, erlang, elixir
function removeTrailingZeros(num: string): string {
  return num.length <= 1 || num[num.length - 1] !== '0'
    ? num
    : removeTrailingZeros(num.slice(0, num.length - 1));
}
