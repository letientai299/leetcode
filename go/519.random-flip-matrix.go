package main

import (
	"math/rand"
)

/*
 * @lc app=leetcode id=519 lang=golang
 *
 * [519] Random Flip Matrix
 *
 * https://leetcode.com/problems/random-flip-matrix/description/
 *
 * algorithms
 * Medium (34.83%)
 * Total Accepted:    9.6K
 * Total Submissions: 25.5K
 * Testcase Example:  '["Solution", "flip", "flip", "flip", "flip"]\n[[2, 2], [], [], [], []]'
 *
 * You are given the number of rows n_rows and number of columns n_cols of a 2D
 * binary matrix where all values are initially 0. Write a function flip which
 * chooses a 0 value uniformly at random, changes it to 1, and then returns the
 * position [row.id, col.id] of that value. Also, write a function reset which
 * sets all values back to 0. Try to minimize the number of calls to system's
 * Math.random() and optimize the time and space complexity.
 *
 * Note:
 *
 *
 * 1 <= n_rows, n_cols <= 10000
 * 0 <= row.id < n_rows and 0 <= col.id < n_cols
 * flip will not be called when the matrix has no 0 values left.
 * the total number of calls to flip and reset will not exceed 1000.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * ["Solution","flip","flip","flip","flip"]
 * [[2,3],[],[],[],[]]
 * Output: [null,[0,1],[1,2],[1,0],[1,1]]
 *
 *
 *
 * Example 2:
 *
 *
 * Input:
 * ["Solution","flip","flip","reset","flip"]
 * [[1,2],[],[],[],[]]
 * Output: [null,[0,0],[0,1],null,[0,0]]
 *
 *
 * Explanation of Input Syntax:
 *
 * The input is two lists: the subroutines called and their arguments.
 * Solution's constructor has two arguments, n_rows and n_cols. flip and reset
 * have no arguments. Arguments are always wrapped with a list, even if there
 * aren't any.
 *
 */
type Solution struct {
	row     int
	col     int
	remain  int
	flipped map[int]int
}

func Constructor(row int, col int) Solution {
	s := Solution{
		col: col,
		row: row,
	}
	s.Reset()
	return s
}

func (s *Solution) Flip() []int {
	r := int(rand.Int31()) % s.remain
	s.remain--
	sel := r
	if val, ok := s.flipped[r]; ok {
		sel = val
	}

	s.flipped[r] = s.remain
	if val, ok := s.flipped[s.remain]; ok {
		s.flipped[r] = val
	}
	return []int{sel / s.col, sel % s.col}
}

func (s *Solution) Reset() {
	s.remain = s.row * s.col
	s.flipped = make(map[int]int)
}
