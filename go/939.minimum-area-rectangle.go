package main

import (
	"sort"
)

// Minimum Area Rectangle
//
// Medium
//
// https://leetcode.com/problems/minimum-area-rectangle/
//
// You are given an array of points in the **X-Y** plane `points` where
// `points[i] = [xi, yi]`.
//
// Return _the minimum area of a rectangle formed from these points, with sides
// parallel to the X and Y axes_. If there is not any such rectangle, return
// `0`.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2021/08/03/rec1.JPG)
//
// ```
// Input: points = [[1,1],[1,3],[3,1],[3,3],[2,2]]
// Output: 4
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2021/08/03/rec2.JPG)
//
// ```
// Input: points = [[1,1],[1,3],[3,1],[3,3],[4,1],[4,3]]
// Output: 2
//
// ```
//
// **Constraints:**
//
// - `1 <= points.length <= 500`
// - `points[i].length == 2`
// - `0 <= xi, yi <= 4 * 104`
// - All the given points are **unique**.
func minAreaRect(points [][]int) int {
	var xs []int
	xmap := make(map[int][]int)
	for id, p := range points {
		x := p[0]
		if xmap[x] == nil {
			xs = append(xs, x)
		}
		xmap[x] = append(xmap[x], id)
	}

	sort.Ints(xs)
	for _, m := range xmap {
		sort.Slice(m, func(i, j int) bool {
			a, b := points[m[i]], points[m[j]]
			return a[1] < b[1]
		})
	}

	best := -1

	for i := 0; i < len(xs)-1; i++ {
		a := xmap[xs[i]]
		if len(a) < 2 {
			continue
		}

		for j := i + 1; j < len(xs); j++ {
			b := xmap[xs[j]]
			if len(b) < 2 {
				continue
			}

			for h, k := 0, 0; h < len(a) && k < len(b); {
				// lower side
				if points[a[h]][1] == points[b[k]][1] {
					for u, v := h+1, k+1; u < len(a) && v < len(b); {
						if points[a[u]][1] == points[b[v]][1] {
							area := (xs[j] - xs[i]) * (points[a[u]][1] - points[a[h]][1])
							if best == -1 || area < best {
								best = area
							}

							u++
							v++
							continue
						}

						if points[a[u]][1] > points[b[v]][1] {
							v++
						} else {
							u++
						}
					}

					h++
					k++
					continue
				}

				if points[a[h]][1] > points[b[k]][1] {
					k++
				} else {
					h++
				}
			}
		}
	}

	if best == -1 {
		best = 0
	}

	return best
}
