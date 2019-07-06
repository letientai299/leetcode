package main

/*
 * @lc app=leetcode id=661 lang=golang
 *
 * [661] Image Smoother
 *
 * https://leetcode.com/problems/image-smoother/description/
 *
 * algorithms
 * Easy (48.95%)
 * Total Accepted:    36.4K
 * Total Submissions: 74.3K
 * Testcase Example:  '[[1,1,1],[1,0,1],[1,1,1]]'
 *
 * Given a 2D integer matrix M representing the gray scale of an image, you
 * need to design a smoother to make the gray scale of each cell becomes the
 * average gray scale (rounding down) of all the 8 surrounding cells and
 * itself.  If a cell has less than 8 surrounding cells, then use as many as
 * you can.
 *
 * Example 1:
 *
 * Input:
 * [[1,1,1],
 * ⁠[1,0,1],
 * ⁠[1,1,1]]
 * Output:
 * [[0, 0, 0],
 * ⁠[0, 0, 0],
 * ⁠[0, 0, 0]]
 * Explanation:
 * For the point (0,0), (0,2), (2,0), (2,2): floor(3/4) = floor(0.75) = 0
 * For the point (0,1), (1,0), (1,2), (2,1): floor(5/6) = floor(0.83333333) = 0
 * For the point (1,1): floor(8/9) = floor(0.88888889) = 0
 *
 *
 *
 * Note:
 *
 * The value in the given matrix is in the range of [0, 255].
 * The length and width of the given matrix are in the range of [1, 150].
 *
 *
 */
func imageSmoother(img [][]int) [][]int {
	// clone the array
	h := len(img)
	if h == 0 {
		return nil
	}

	w := len(img[0])

	newImg := make([][]int, h)
	for y := 0; y < h; y++ {
		rows := make([]int, w)
		for x := 0; x < w; x++ {
			pixels := 1
			sum := img[y][x]

			if y+1 < h {
				pixels++
				sum += img[y+1][x]
			}

			if y-1 >= 0 {
				pixels++
				sum += img[y-1][x]
			}

			if x+1 < w {
				pixels++
				sum += img[y][x+1]
			}

			if x-1 >= 0 {
				pixels++
				sum += img[y][x-1]
			}

			if y+1 < h && x+1 < w {
				pixels++
				sum += img[y+1][x+1]
			}

			if y+1 < h && x-1 >= 0 {
				pixels++
				sum += img[y+1][x-1]
			}

			if y-1 >= 0 && x-1 >= 0 {
				pixels++
				sum += img[y-1][x-1]
			}

			if y-1 >= 0 && x+1 < w {
				pixels++
				sum += img[y-1][x+1]
			}

			rows[x] = sum / pixels
		}
		newImg[y] = rows
	}

	return newImg
}
