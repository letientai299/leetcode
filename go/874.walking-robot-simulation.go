package main

/*
 * @lc app=leetcode id=874 lang=golang
 *
 * [874] Walking Robot Simulation
 *
 * https://leetcode.com/problems/walking-robot-simulation/description/
 *
 * algorithms
 * Easy (33.51%)
 * Total Accepted:    21.8K
 * Total Submissions: 59.7K
 * Testcase Example:  '[4,-1,3]\n[]'
 *
 * A robot on an infinite grid starts at point (0, 0) and faces north.  The
 * robot can receive one of three possible types of commands:
 *
 *
 * -2: turn left 90 degrees
 * -1: turn right 90 degrees
 * 1 <= x <= 9: move forward x units
 *
 *
 * Some of the grid squares are obstacles.
 *
 * The i-th obstacle is at grid point (obstacles[i][0], obstacles[i][1])
 *
 * If the robot would try to move onto them, the robot stays on the previous
 * grid square instead (but still continues following the rest of the route.)
 *
 * Return the square of the maximum Euclidean distance that the robot will be
 * from the origin.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: commands = [4,-1,3], obstacles = []
 * Output: 25
 * Explanation: robot will go to (3, 4)
 *
 *
 *
 * Example 2:
 *
 *
 * Input: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
 * Output: 65
 * Explanation: robot will be stuck at (1, 4) before turning left and going to
 * (1, 8)
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 0 <= commands.length <= 10000
 * 0 <= obstacles.length <= 10000
 * -30000 <= obstacle[i][0] <= 30000
 * -30000 <= obstacle[i][1] <= 30000
 * The answer is guaranteed to be less than 2 ^ 31.
 *
 *
 */

func robotSim(commands []int, obstacles [][]int) int {
	x, y := 0, 0
	dx, dy := 0, 1
	m := make(map[int]map[int]bool)
	for _, a := range obstacles {
		if m[a[0]] == nil {
			m[a[0]] = make(map[int]bool)
		}

		m[a[0]][a[1]] = true
	}

	dist := 0
	for _, c := range commands {
		switch c {
		case -2:
			dx, dy = -dy, dx
		case -1:
			dx, dy = dy, -dx
		default:
			for c > 0 && !m[x+dx][y+dy] {
				c--
				x += dx
				y += dy
			}
			d := x*x + y*y
			if dist < d {
				dist = d
			}
		}
	}
	return dist
}
