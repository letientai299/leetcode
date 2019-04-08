package main

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

func calPointDistance(p1, p2 []int) int {

	x := p2[0] - p1[0]
	y := p2[1] - p1[1]
	return x*x + y*y
}

func numberOfBoomerangs(points [][]int) int {
	var count int
	for i, p := range points {
		distance := make(map[int]int, len(points))
		for j, p1 := range points {
			if i != j {
				distance[calPointDistance(p1, p)]++
			}
		}
		for _, v := range distance {
			count = count + v*(v-1)
		}
	}
	return count
}
