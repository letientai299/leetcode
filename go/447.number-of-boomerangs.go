package main

import (
	"fmt"
	"math"
)

/*
* @lc app=leetcode id=447 lang=golang
*
* [447] Number of Boomerangs
*
* https://leetcode.com/problems/number-of-boomerangs/description/
*
* algorithms
* Easy (49.55%)
* Total Accepted:    52.4K
* Total Submissions: 105.7K
* Testcase Example:  '[[0,0],[1,0],[2,0]]'
*
* Given n points in the plane that are all pairwise distinct, a "boomerang" is
* a tuple of points (i, j, k) such that the distance between i and j equals
* the distance between i and k (the order of the tuple matters).
*
* Find the number of boomerangs. You may assume that n will be at most 500 and
* coordinates of points are all in the range [-10000, 10000] (inclusive).
*
* Example:
*
* Input:
* [[0,0],[1,0],[2,0]]
*
* Output:
* 2
*
* Explanation:
* The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]]
*
*
 */
func numberOfBoomerangs(points [][]int) int {
	count := 0
	for i, a := range points {
		for j, b := range points {
			if i >= j {
				continue
			}

			x0 := 0.5 * float64(a[0]+b[0])
			y0 := 0.5 * float64(a[1]+b[1])
			v := float64(b[1] - a[1])
			w := float64(a[0] - b[0])

			f := func(x, y float64) bool {
				if v == 0 {
					return x == x0
				}

				if w == 0 {
					return y == y0
				}
				return 0.0000000001 > math.Abs(w*(x-x0) - v*(y-y0))
			}

			for k, c := range points {
				if k == j || k == i {
					continue
				}

				if f(float64(c[0]), float64(c[1])) {
					fmt.Printf("(%d, %d, %d)\n", k, i, j)
					fmt.Printf("(%d, %d, %d)\n", k, j, i)
					count += 2
				}
			}
		}
	}
	return count
}
