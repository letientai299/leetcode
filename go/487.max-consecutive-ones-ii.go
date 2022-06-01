package main

// Max Consecutive Ones II
//
// Medium
//
// https://leetcode.com/problems/max-consecutive-ones-ii/
func findMaxConsecutiveOnes(nums []int) int {
	best := 0
	flipped := 0
	unflip := 0
	for _, n := range nums {
		if n == 1 {
			unflip++
			flipped++
			continue
		}

		if best < flipped {
			best = flipped
		}

		flipped = unflip + 1
		unflip = 0
	}

	if best < flipped {
		best = flipped
	}

	return best
}
