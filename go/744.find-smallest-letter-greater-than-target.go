package main

/*
 * @lc app=leetcode id=744 lang=golang
 *
 * [744] Find Smallest Letter Greater Than Target
 *
 * https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/
 *
 * algorithms
 * Easy (44.23%)
 * Total Accepted:    45.1K
 * Total Submissions: 101.8K
 * Testcase Example:  '["c","f","j"]\n"a"'
 *
 *
 * Given a list of sorted characters letters containing only lowercase letters,
 * and given a target letter target, find the smallest element in the list that
 * is larger than the given target.
 *
 * Letters also wrap around.  For example, if the target is target = 'z' and
 * letters = ['a', 'b'], the answer is 'a'.
 *
 *
 * Examples:
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "a"
 * Output: "c"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "c"
 * Output: "f"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "d"
 * Output: "f"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "g"
 * Output: "j"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "j"
 * Output: "c"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "k"
 * Output: "c"
 *
 *
 *
 * Note:
 *
 * letters has a length in range [2, 10000].
 * letters consists of lowercase letters, and contains at least 2 unique
 * letters.
 * target is a lowercase letter.
 *
 *
 */
func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	nums := make([]int, n)

	// convert letter to int and remove the character rotation condition
	prev := -1
	for i, c := range letters {
		nums[i] = int(c - 'a')
		if nums[i] < prev {
			nums[i] += 'z' - 'a'
		}
		prev = nums[i]
	}

	x := int(target - 'a')
	// do a binary search to look for target
	left, right := 0, n
	found := -1
	for left < right {
		mid := (right + left) / 2
		if nums[mid] == x {
			found = mid
			break
		}

		if nums[mid] > x {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if found < 0 { // not found
		found = left
	}

	for found < n && nums[found] <= x {
		found++
	}

	if found >= n {
		return letters[0]
	}

	if nums[found] > x {
		return letters[found]
	}

	if found+1 >= n {
		return letters[0]
	}

	return letters[found+1]
}
