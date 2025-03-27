package main

// Count Vowel Substrings of a String
//
// Easy
//
// https://leetcode.com/problems/count-vowel-substrings-of-a-string
//
// A **substring** is a contiguous (non-empty) sequence of characters within a
// string.
//
// A **vowel substring** is a substring that **only** consists of vowels (`'a'`,
// `'e'`, `'i'`, `'o'`, and `'u'`) and has **all five** vowels present in it.
//
// Given a string `word`, return _the number of **vowel substrings** in_ `word`.
//
// **Example 1:**
//
// ```
// Input: word = "aeiouu"
// Output: 2
// Explanation: The vowel substrings of word are as follows (underlined):
// - "aeiouu"
// - "aeiouu"
//
// ```
//
// **Example 2:**
//
// ```
// Input: word = "unicornarihan"
// Output: 0
// Explanation: Not all 5 vowels are present, so there are no vowel substrings.
//
// ```
//
// **Example 3:**
//
// ```
// Input: word = "cuaieuouac"
// Output: 7
// Explanation: The vowel substrings of word are as follows (underlined):
// - "cuaieuouac"
// - "cuaieuouac"
// - "cuaieuouac"
// - "cuaieuouac"
// - "cuaieuouac"
// - "cuaieuouac"
// - "cuaieuouac"
//
// ```
//
// **Constraints:**
//
// - `1 <= word.length <= 100`
// - `word` consists of lowercase English letters only.
func countVowelSubstrings(s string) int {
  n := len(s)
  res := 0
  for l := 0; l < n; l++ {
    // find substring that only contains vowels
    r := l
    for r < n && isVowel(s[r]) {
      r++
    }

    res += _countVowelSubstrings(s[l:r])
    l = r
  }

  return res
}

func _countVowelSubstrings(s string) int {
  if len(s) < 5 {
    return 0
  }

  n := len(s)
  res := 0
  for l := 0; l < n; l++ {
    vowels := make(map[byte]int, 5)
    for r := l; r < n; r++ {
      vowels[s[r]]++
      if len(vowels) == 5 {
        res++
      }
    }
  }

  return res
}

// TODO (tai): can be faster, best is 0ms, I got 3ms
