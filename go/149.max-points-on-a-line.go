package main

import (
	"fmt"
	"math"
)

// Max Points on a Line
//
// Hard
//
// https://leetcode.com/problems/max-points-on-a-line/
//
// Given an array of `points` where `points[i] = [xi, yi]` represents a point on
// the **X-Y** plane, return _the maximum number of points that lie on the same
// straight line_.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2021/02/25/plane1.jpg)
//
// ```
// Input: points = [[1,1],[2,2],[3,3]]
// Output: 3
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2021/02/25/plane2.jpg)
//
// ```
// Input: points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
// Output: 4
//
// ```
//
// **Constraints:**
//
// - `1 <= points.length <= 300`
// - `points[i].length == 2`
// - `-104 <= xi, yi <= 104`
// - All the `points` are **unique**.
func maxPoints(points [][]int) int {
	// Damn! Just use brute force would be faster given the input size.

	vectors := make(map[string]int)
	xs := make(map[int]int)
	ys := make(map[int]int)

	for i, a := range points {
		for _, b := range points[i+1:] {
			if a[1] == b[1] {
				// same y
				ys[a[1]]++
			} else if a[0] == b[0] {
				// same x
				xs[a[0]]++
			} else {
				// calculate the line vector, but we can't use float64 due to its preciseness
				// hence, we will store 4 values: dx, dy and m/n (value where the line
				// hit Oy)
				dx, dy := a[0]-b[0], a[1]-b[1]
				if dy < 0 {
					dx *= -1
					dy *= -1
				}
				g := gcd(abs(dx), abs(dy))
				dx /= g
				dy /= g

				m := a[1]*dx - a[0]*dy
				vectors[fmt.Sprintf("%d,%d,%d", dx, dy, m)]++
			}
		}
	}

	best := 1
	for _, v := range vectors {
		x := int(math.Floor(math.Sqrt(float64(v*2)))) + 1
		if x > best {
			best = x
		}
	}

	for _, v := range xs {
		x := int(math.Floor(math.Sqrt(float64(v*2)))) + 1
		if x > best {
			best = x
		}
	}

	for _, v := range ys {
		x := int(math.Floor(math.Sqrt(float64(v*2)))) + 1
		if x > best {
			best = x
		}
	}

	return best
}
