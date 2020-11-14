package main

/*
 * @lc app=leetcode id=925 lang=golang
 *
 * [925] Long Pressed Name
 *
 * https://leetcode.com/problems/long-pressed-name/description/
 *
 * algorithms
 * Easy (44.34%)
 * Total Accepted:    53.9K
 * Total Submissions: 138.1K
 * Testcase Example:  '"alex"\n"aaleex"'
 *
 * Your friend is typing his name into a keyboard.  Sometimes, when typing a
 * character c, the key might get long pressed, and the character will be typed
 * 1 or more times.
 *
 * You examine the typed characters of the keyboard.  Return True if it is
 * possible that it was your friends name, with some characters (possibly none)
 * being long pressed.
 *
 *
 * Example 1:
 *
 *
 * Input: name = "alex", typed = "aaleex"
 * Output: true
 * Explanation: 'a' and 'e' in 'alex' were long pressed.
 *
 *
 * Example 2:
 *
 *
 * Input: name = "saeed", typed = "ssaaedd"
 * Output: false
 * Explanation: 'e' must have been pressed twice, but it wasn't in the typed
 * output.
 *
 *
 * Example 3:
 *
 *
 * Input: name = "leelee", typed = "lleeelee"
 * Output: true
 *
 *
 * Example 4:
 *
 *
 * Input: name = "laiden", typed = "laiden"
 * Output: true
 * Explanation: It's not necessary to long press any character.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= name.length <= 1000
 * 1 <= typed.length <= 1000
 * The characters of name and typed are lowercase letters.
 *
 *
 */
func isLongPressedName(name string, typed string) bool {
	x, y := 0, 0
	for ; x < len(name) && y < len(typed); x++ {
		if name[x] != typed[y] {
			return false
		}

		for x+1 < len(name) && y+1 < len(typed) {
			if name[x+1] == name[x] && typed[y+1] == typed[y] {
				x++
				y++
			} else {
				break
			}
		}

		for y < len(typed) && name[x] == typed[y] {
			y++
		}
	}
	return x == len(name) && y == len(typed)
}
