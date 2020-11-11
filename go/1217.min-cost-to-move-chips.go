package main

// https://leetcode.com/problems/minimum-cost-to-move-chips-to-the-same-position/
//
// Just math
func minCostToMoveChips(position []int) int {
	odd := 0
	even := 0
	for _, x := range position {
		if x%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	if even > odd {
		return odd
	}

	return even
}
