package main

/*
* @lc app=leetcode id=405 lang=golang
*
* [405] Convert a Number to Hexadecimal
*
* https://leetcode.com/problems/convert-a-number-to-hexadecimal/description/
*
* algorithms
* Easy (41.74%)
* Total Accepted:    45K
* Total Submissions: 107.8K
* Testcase Example:  '26'
*
*
* Given an integer, write an algorithm to convert it to hexadecimal. For
* negative integer, two’s complement method is used.
*
*
* Note:
*
* All letters in hexadecimal (a-f) must be in lowercase.
* The hexadecimal string must not contain extra leading 0s. If the number is
* zero, it is represented by a single zero character '0'; otherwise, the first
* character in the hexadecimal string will not be the zero character.
* The given number is guaranteed to fit within the range of a 32-bit signed
* integer.
* You must not use any method provided by the library which converts/formats
* the number to hex directly.
*
*
*
* Example 1:
*
* Input:
* 26
*
* Output:
* "1a"
*
*
*
* Example 2:
*
* Input:
* -1
*
* Output:
* "ffffffff"
*
*
 */
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	// 2 complement example:
	// if a + b == 10000 (in binary)
	// then a and b is 2 complements of each others under 2^4
	if num < 0 {
		// here's 2 complement calculation for 32 bit numbers
		num = int(int64(1<<32) + int64(num))
	}

	base := "0123456789abcdef"
	s := make([]byte, 8)
	for i := 0; i < 8; i++ {
		if num == 0 {
			return string(s[8-i:])
		}
		s[7-i] = base[num%16]
		num >>= 4
	}

	return string(s)
}

// -3: _0000
// -2: 01110
// -1: 01111
//  0: _0000
//  1: _0001
//  2: _0010
//  3: _0011
//  4: _0100
//  5: _0101
//  6: _0110
//  7: _0111
