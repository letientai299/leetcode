package main

// Decode the Message
//
// # Easy
//
// https://leetcode.com/problems/decode-the-message
//
// You are given the strings `key` and `message`, which represent a cipher key
// and a secret message, respectively. The steps to decode `message` are as
// follows:
//
// 1. Use the **first** appearance of all 26 lowercase English letters in `key`
// as the **order** of the substitution table.
// 2. Align the substitution table with the regular English alphabet.
// 3. Each letter in `message` is then **substituted** using the table.
// 4. Spaces `' '` are transformed to themselves.
//
// - For example, given `key = "happy boy"` (actual key would have **at least
// one** instance of each letter in the alphabet), we have the partial
// substitution table of (`'h' -> 'a'`, `'a' -> 'b'`, `'p' -> 'c'`, `'y' ->
// 'd'`, `'b' -> 'e'`, `'o' -> 'f'`).
//
// Return _the decoded message_.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2022/05/08/ex1new4.jpg)
//
// ```
// Input: key = "the quick brown fox jumps over the lazy dog", message = "vkbs
// bs t suepuv"
// Output: "this is a secret"
// Explanation: The diagram above shows the substitution table.
// It is obtained by taking the first appearance of each letter in "the quick
// brown fox jumps over the lazy dog".
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2022/05/08/ex2new.jpg)
//
// ```
// Input: key = "eljuxhpwnyrdgtqkviszcfmabo", message = "zwx hnfx lqantp mnoeius
// ycgk vcnjrdb"
// Output: "the five boxing wizards jump quickly"
// Explanation: The diagram above shows the substitution table.
// It is obtained by taking the first appearance of each letter in
// "eljuxhpwnyrdgtqkviszcfmabo".
//
// ```
//
// **Constraints:**
//
// - `26 <= key.length <= 2000`
// - `key` consists of lowercase English letters and `' '`.
// - `key` contains every letter in the English alphabet (`'a'` to `'z'`) **at
// least once**.
// - `1 <= message.length <= 2000`
// - `message` consists of lowercase English letters and `' '`.
func decodeMessage(key string, message string) string {
	var tbl [26]uint8
	seen := uint(0)
	ti := uint8(0)
	for _, c := range []byte(key) {
		if c == ' ' {
			continue
		}
		idx := c - 'a' // letter index
		bit := uint(1 << idx)
		//fmt.Printf("before %32s   %d\n", strconv.FormatInt(int64(seen), 2), seen)
		//fmt.Printf(" bit   %32s   %d\n", strconv.FormatInt(int64(bit), 2), bit)
		//fmt.Printf(" test  %32s\n", strconv.FormatInt(int64(seen&bit), 2))
		if seen&bit == 0 {
			seen |= bit
			//fmt.Printf(" after %32s\n", strconv.FormatInt(int64(seen), 2))
			tbl[idx] = ti
			ti++
		}
	}

	bs := []byte(message)
	for i, c := range bs {
		if c != ' ' {
			bs[i] = tbl[c-'a'] + 'a'
		}
	}

	return string(bs)
}
