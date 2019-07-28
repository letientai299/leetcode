package main

/*
 * @lc app=leetcode id=733 lang=golang
 *
 * [733] Flood Fill
 *
 * https://leetcode.com/problems/flood-fill/description/
 *
 * algorithms
 * Easy (51.19%)
 * Total Accepted:    53.9K
 * Total Submissions: 105.1K
 * Testcase Example:  '[[1,1,1],[1,1,0],[1,0,1]]\n1\n1\n2'
 *
 *
 * An image is represented by a 2-D array of integers, each integer
 * representing the pixel value of the image (from 0 to 65535).
 *
 * Given a coordinate (sr, sc) representing the starting pixel (row and column)
 * of the flood fill, and a pixel value newColor, "flood fill" the image.
 *
 * To perform a "flood fill", consider the starting pixel, plus any pixels
 * connected 4-directionally to the starting pixel of the same color as the
 * starting pixel, plus any pixels connected 4-directionally to those pixels
 * (also with the same color as the starting pixel), and so on.  Replace the
 * color of all of the aforementioned pixels with the newColor.
 *
 * At the end, return the modified image.
 *
 * Example 1:
 *
 * Input:
 * image = [[1,1,1],
 *          [1,1,0],
 *          [1,0,1]]
 * sr = 1, sc = 1, newColor = 2
 * Output: [[2,2,2],
 *          [2,2,0],
 *          [2,0,1]]
 * Explanation:
 * From the center of the image (with position (sr, sc) = (1, 1)), all pixels
 * connected
 * by a path of the same color as the starting pixel are colored with the new
 * color.
 * Note the bottom corner is not colored 2, because it is not 4-directionally
 * connected
 * to the starting pixel.
 *
 *
 *
 * Note:
 * The length of image and image[0] will be in the range [1, 50].
 * The given starting pixel will satisfy 0  and 0 .
 * The value of each color in image[i][j] and newColor will be an integer in
 * [0, 65535].
 *
 */
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	m := len(image)
	n := len(image[0])
	checks := make([][]bool, m)
	for i := range checks {
		checks[i] = make([]bool, n)
	}

	cur := image[sr][sc]

	var fill func(x, y int)
	fill = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}

		if image[x][y] != cur {
			return
		}

		if !checks[x][y] {
			checks[x][y] = true
		} else {
			return
		}

		image[x][y] = newColor
		fill(x-1, y)
		fill(x+1, y)
		fill(x, y-1)
		fill(x, y+1)
	}

	fill(sr, sc)
	return image
}
