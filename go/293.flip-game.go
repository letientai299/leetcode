package main

/*
 * @lc app=leetcode id=293 lang=golang
 *
 * [293] Flip Game
 *
 * https://leetcode.com/problems/flip-game/description/
 *
 * algorithms
 * Easy (59.63%)
 * Total Accepted:    50.9K
 * Total Submissions: 83.6K
 * Testcase Example:  '"++++"'
 *
 * You are playing the following Flip Game with your friend: Given a string
 * that contains only these two characters: + and -, you and your friend take
 * turns to flip two consecutive "++" into "--". The game ends when a person
 * can no longer make a move and therefore the other person will be the
 * winner.
 *
 * Write a function to compute all possible states of the string after one
 * valid move.
 *
 * Example:
 *
 *
 * Input: s = "++++"
 * Output:
 * [
 * ⁠ "--++",
 * ⁠ "+--+",
 * ⁠ "++--"
 * ]
 *
 *
 * Note: If there is no valid move, return an empty list [].
 *
 */
func generatePossibleNextMoves(s string) []string {
	var res []string
	for i := 0; i < len(s); {
		if s[i] == '-' {
			i++
			continue
		}

		if i >= len(s)-1 || s[i+1] == '-' {
			i += 2
			continue
		}

		var pre string
		if i-1 >= 0 {
			pre = s[:i]
		}
		var post string
		if i+2 < len(s) {
			post = s[i+2:]
		}
		res = append(res, pre+"--"+post)
		i++
	}

	return res
}
