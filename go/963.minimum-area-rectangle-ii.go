package main

import "math"

/*
 * @lc app=leetcode id=963 lang=golang
 *
 * [963] Minimum Area Rectangle II
 *
 * https://leetcode.com/problems/minimum-area-rectangle-ii/description/
 *
 * algorithms
 * Medium (48.05%)
 * Total Accepted:    13.9K
 * Total Submissions: 27K
 * Testcase Example:  '[[1,2],[2,1],[1,0],[0,1]]'
 *
 * Given a set of points in the xy-plane, determine the minimum area of any
 * rectangle formed from these points, with sides not necessarily parallel to
 * the x and y axes.
 *
 * If there isn't any rectangle, return 0.
 *
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: [[1,2],[2,1],[1,0],[0,1]]
 * Output: 2.00000
 * Explanation: The minimum area rectangle occurs at [1,2],[2,1],[1,0],[0,1],
 * with an area of 2.
 *
 *
 *
 * Example 2:
 *
 *
 *
 *
 * Input: [[0,1],[2,1],[1,1],[1,0],[2,0]]
 * Output: 1.00000
 * Explanation: The minimum area rectangle occurs at [1,0],[1,1],[2,1],[2,0],
 * with an area of 1.
 *
 *
 *
 * Example 3:
 *
 *
 *
 *
 * Input: [[0,3],[1,2],[3,1],[1,3],[2,1]]
 * Output: 0
 * Explanation: There is no possible rectangle to form from these points.
 *
 *
 *
 * Example 4:
 *
 *
 *
 *
 * Input: [[3,1],[1,1],[0,1],[2,1],[3,3],[3,2],[0,2],[2,3]]
 * Output: 2.00000
 * Explanation: The minimum area rectangle occurs at [2,1],[2,3],[3,3],[3,1],
 * with an area of 2.
 *
 *
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= points.length <= 50
 * 0 <= points[i][0] <= 40000
 * 0 <= points[i][1] <= 40000
 * All points are distinct.
 * Answers within 10^-5 of the actual value will be accepted as correct.
 *
 *
 */

/*
func minAreaFreeRect(points [][]int) float64 {
	type point [2]int
	type centerAndDistance struct {
		x, y float64
		d    int
	}

	n := len(points)

	// Previous implement use 2 level maps, got 28ms, 33%
	//    center -> {distance -> point_pairs}
	// Switch to single level map like this, got 16ms, 66%
	m := make(map[centerAndDistance][][2]point, n^n)

	area := func(x, y, z point) float64 {
		long := (x[0]-y[0])*(x[0]-y[0]) + (x[1]-y[1])*(x[1]-y[1])
		short := (x[0]-z[0])*(x[0]-z[0]) + (x[1]-z[1])*(x[1]-z[1])
		return math.Sqrt(float64(long * short))
	}

	smallest := float64(0)

	for i := 0; i < n-1; i++ {
		x := points[i]
		for j := i + 1; j < n; j++ {
			y := points[j]
			c := centerAndDistance{
				x: float64(x[0]+y[0]) / 2,
				y: float64(x[1]+y[1]) / 2,
				d: (x[0]-y[0])*(x[0]-y[0]) + (x[1]-y[1])*(x[1]-y[1]),
			}

			a := [2]point{
				{x[0], x[1]},
				{y[0], y[1]},
			}

			if len(m[c]) > 0 {
				for _, b := range m[c] {
					s := area(a[0], b[0], b[1])
					if smallest == 0 || s < smallest {
						smallest = s
					}
				}
			}

			m[c] = append(m[c], a)
		}
	}

	return smallest
}
*/

func minAreaFreeRect(points [][]int) float64 {
	type point [2]int
	n := len(points)
	set := make(map[point]struct{}, n)

	for _, p := range points {
		set[point{p[0], p[1]}] = struct{}{}
	}

	smallest := float64(0)

	for i := 0; i < n-2; i++ {
		x := points[i]
		for j := i + 1; j < n-1; j++ {
			y := points[j]
			for k := j + 1; k < n; k++ {
				z := points[k]

				// this check is much more faster than the check for equal diagonal lines
				// From 20ms to 4ms
				if !isOrthogonal([]int{x[0] - z[0], x[1] - z[1]}, []int{y[0] - z[0], y[1] - z[1]}) {
					continue
				}

				t := point{x[0] + y[0] - z[0], x[1] + y[1] - z[1]}
				if _, ok := set[t]; !ok {
					continue
				}

				s := area(z, x, y)
				if smallest == 0 || s < smallest {
					smallest = s
				}
			}
		}
	}

	return smallest
}

func area(x, y, z []int) float64 {
	long := (x[0]-y[0])*(x[0]-y[0]) + (x[1]-y[1])*(x[1]-y[1])
	short := (x[0]-z[0])*(x[0]-z[0]) + (x[1]-z[1])*(x[1]-z[1])
	return math.Sqrt(float64(long * short))
}

func isOrthogonal(a, b []int) bool {
	return a[0]*b[0]+a[1]*b[1] == 0
}
