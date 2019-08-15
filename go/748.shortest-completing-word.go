package main

/*
 * @lc app=leetcode id=748 lang=golang
 *
 * [748] Shortest Completing Word
 *
 * https://leetcode.com/problems/shortest-completing-word/description/
 *
 * algorithms
 * Easy (54.56%)
 * Total Accepted:    22.2K
 * Total Submissions: 40.5K
 * Testcase Example:  '"1s3 PSt"\n["step","steps","stripe","stepple"]'
 *
 *
 * Find the minimum length word from a given dictionary words, which has all
 * the letters from the string licensePlate.  Such a word is said to complete
 * the given string licensePlate
 *
 * Here, for letters we ignore case.  For example, "P" on the licensePlate
 * still matches "p" on the word.
 *
 * It is guaranteed an answer exists.  If there are multiple answers, return
 * the one that occurs first in the array.
 *
 * The license plate might have the same letter occurring multiple times.  For
 * example, given a licensePlate of "PP", the word "pair" does not complete the
 * licensePlate, but the word "supper" does.
 *
 *
 * Example 1:
 *
 * Input: licensePlate = "1s3 PSt", words = ["step", "steps", "stripe",
 * "stepple"]
 * Output: "steps"
 * Explanation: The smallest length word that contains the letters "S", "P",
 * "S", and "T".
 * Note that the answer is not "step", because the letter "s" must occur in the
 * word twice.
 * Also note that we ignored case for the purposes of comparing whether a
 * letter exists in the word.
 *
 *
 *
 * Example 2:
 *
 * Input: licensePlate = "1s3 456", words = ["looks", "pest", "stew", "show"]
 * Output: "pest"
 * Explanation: There are 3 smallest length words that contains the letters
 * "s".
 * We return the one that occurred first.
 *
 *
 *
 * Note:
 *
 * licensePlate will be a string with length in range [1, 7].
 * licensePlate will contain digits, spaces, or letters (uppercase or
 * lowercase).
 * words will have a length in the range [10, 1000].
 * Every words[i] will consist of lowercase letters, and have length in range
 * [1, 15].
 *
 *
 */
func shortestCompletingWord(licensePlate string, words []string) string {
	chars := make([]int, 26)

	// count for the license plate
	for _, c := range licensePlate {
		if 'A' <= c && 'Z' >= c {
			c = c + 'a' - 'A'
		}

		if 'a' <= c && 'z' >= c {
			chars[c-'a']++
		}
	}

	var shorted string
	for _, word := range words {
		ws := make([]int, 26)
		for _, w := range word {
			if 'A' <= w && 'Z' >= w {
				w = w + 'a' - 'A'
			}

			if 'a' > w || 'z' < w {
				continue
			}

			ws[w-'a'] ++
		}

		ok := true
		for i := range ws {
			if ws[i] < chars[i] {
				ok = false
				break
			}
		}
		if ok && (shorted == "" || len(shorted) > len(word)) {
			shorted = word
		}
	}

	return shorted
}
