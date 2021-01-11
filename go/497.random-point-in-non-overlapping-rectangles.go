package main

import (
	"math/rand"
	"sort"
)

/*
 * @lc app=leetcode id=497 lang=golang
 *
 * [497] Random Point in Non-overlapping Rectangles
 *
 * https://leetcode.com/problems/random-point-in-non-overlapping-rectangles/description/
 *
 * algorithms
 * Medium (36.93%)
 * Total Accepted:    29.9K
 * Total Submissions: 76.5K
 * Testcase Example:  '["Solution", "pick", "pick", "pick"]\n[[[[1, 1, 5, 5]]], [], [], []]'
 *
 * Given a list of non-overlapping axis-aligned rectangles rects, write a
 * function pick which randomly and uniformily picks an integer point in the
 * space covered by the rectangles.
 *
 * Note:
 *
 *
 * An integer point is a point that has integer coordinates.
 * A point on the perimeter of a rectangle is included in the space covered by
 * the rectangles.
 * ith rectangle = rects[i] = [x1,y1,x2,y2], where [x1, y1] are the integer
 * coordinates of the bottom-left corner, and [x2, y2] are the integer
 * coordinates of the top-right corner.
 * length and width of each rectangle does not exceed 2000.
 * 1 <= rects.length <= 100
 * pick return a point as an array of integer coordinates [p_x, p_y]
 * pick is called at most 10000 times.
 *
 *
 *
 * Example 1:
 *
 *
 * Input:
 * ["Solution","pick","pick","pick"]
 * [[[[1,1,5,5]]],[],[],[]]
 * Output:
 * [null,[4,1],[4,1],[3,3]]
 *
 *
 *
 * Example 2:
 *
 *
 * Input:
 * ["Solution","pick","pick","pick","pick","pick"]
 * [[[[-2,-2,-1,-1],[1,0,3,0]]],[],[],[],[],[]]
 * Output:
 * [null,[-1,-2],[2,0],[-2,-1],[3,0],[-2,-2]]
 *
 *
 *
 * Explanation of Input Syntax:
 *
 * The input is two lists: the subroutines called and their arguments.
 * Solution's constructor has one argument, the array of rectangles rects. pick
 * has no arguments. Arguments are always wrapped with a list, even if there
 * aren't any.
 *
 *
 *
 *
 *
 *
 *
 */

type Solution497 struct {
	rects   [][]int
	weights []int
	all     int
}

func Constructor497(rects [][]int) Solution497 {
	s := Solution497{
		rects:   rects,
		weights: make([]int, 0, len(rects)),
	}
	w := 0
	for _, r := range rects {
		w += (r[2] - r[0] + 1) * (r[3] - r[1] + 1)
		s.weights = append(s.weights, w)
	}
	s.all = w
	return s
}

func (s *Solution497) Pick() []int {
	w := rand.Intn(s.all) + 1
	i := sort.SearchInts(s.weights, w)
	if i > 0 {
		w -= s.weights[i-1]
	}
	return s.pickRect(i, w)
}

func (s *Solution497) pickRect(i int, w int) []int {
	r := s.rects[i]
	w--
	dx := r[2] - r[0] + 1
	return []int{w%dx + r[0], w/dx + r[1]}
}
