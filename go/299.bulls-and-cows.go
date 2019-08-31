package main

import "fmt"

/*
 * @lc app=leetcode id=299 lang=golang
 *
 * [299] Bulls and Cows
 *
 * https://leetcode.com/problems/bulls-and-cows/description/
 *
 * algorithms
 * Easy (40.30%)
 * Total Accepted:    106.7K
 * Total Submissions: 264.9K
 * Testcase Example:  '"1807"\n"7810"'
 *
 * You are playing the following Bulls and Cows game with your friend: You
 * write down a number and ask your friend to guess what the number is. Each
 * time your friend makes a guess, you provide a hint that indicates how many
 * digits in said guess match your secret number exactly in both digit and
 * position (called "bulls") and how many digits match the secret number but
 * locate in the wrong position (called "cows"). Your friend will use
 * successive guesses and hints to eventually derive the secret number.
 *
 * Write a function to return a hint according to the secret number and
 * friend's guess, use A to indicate the bulls and B to indicate the cows.
 *
 * Please note that both secret number and friend's guess may contain duplicate
 * digits.
 *
 * Example 1:
 *
 *
 * Input: secret = "1807", guess = "7810"
 *
 * Output: "1A3B"
 *
 * Explanation: 1 bull and 3 cows. The bull is 8, the cows are 0, 1 and 7.
 *
 * Example 2:
 *
 *
 * Input: secret = "1123", guess = "0111"
 *
 * Output: "1A1B"
 *
 * Explanation: The 1st 1 in friend's guess is a bull, the 2nd or 3rd 1 is a
 * cow.
 *
 * Note: You may assume that the secret number and your friend's guess only
 * contain digits, and their lengths are always equal.
 */
func getHint(secret string, guess string) string {
	type set map[int]struct{}
	statSecret := make(map[int32]set)
	for i, c := range secret {
		if statSecret[c] == nil {
			statSecret[c] = set{}
		}
		statSecret[c][i] = struct{}{}
	}

	statGuess := make(map[int32]set)
	for i, c := range guess {
		if statGuess[c] == nil {
			statGuess[c] = set{}
		}
		statGuess[c][i] = struct{}{}
	}

	var bulls int
	var cows int
	for c, guessIndexes := range statGuess {
		secretIndexes, ok := statSecret[c]
		if !ok {
			continue
		}

		for i := range secretIndexes {
			if _, ok := guessIndexes[i]; ok {
				bulls++
			} else {
				cows++
			}
		}

		if len(secretIndexes) > len(guessIndexes) {
			cows -= len(secretIndexes) - len(guessIndexes)
		}

	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}
