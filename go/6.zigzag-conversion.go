package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 *
 * https://leetcode.com/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (33.59%)
 * Total Accepted:    382K
 * Total Submissions: 1.1M
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
 * of rows like this: (you may want to display this pattern in a fixed font for
 * better legibility)
 *
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 *
 * And then read line by line: "PAHNAPLSIIGYIR"
 *
 * Write the code that will take a string and make this conversion given a
 * number of rows:
 *
 *
 * string convert(string s, int numRows);
 *
 * Example 1:
 *
 *
 * Input: s = "PAYPALISHIRING", numRows = 3
 * Output: "PAHNAPLSIIGYIR"
 *
 *
 * Example 2:
 *
 *
 * Input: s = "PAYPALISHIRING", numRows = 4
 * Output: "PINALSIGYAHRPI"
 * Explanation:
 *
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 *
 */
func convert(s string, k int) string {
	if k >= len(s) || k < 2 {
		return s
	}

	// build the 2D char array
	cs := make([][]byte, 0)

	n := len(s)
	i := 0
	j := 0
	for i < n {
		end := i + k
		if end > n {
			end = n
		}
		line := make([]byte, k)
		for j = i; j < end; j++ {
			line[j-i] = s[j]
		}
		cs = append(cs, line)

		i = end
		end += k - 2
		if end > n {
			end = n
		}

		line = make([]byte, k)
		// reverse the diagonal line
		for j = i; j < end; j++ {
			line[k-2-j+i] = s[j]
		}
		cs = append(cs, line)

		i = end
	}

	var sb strings.Builder
	for i = 0; i < k; i++ {
		for j = 0; j < len(cs); j++ {
			c := cs[j][i]
			if c > 0 {
				sb.WriteByte(c)
			}
		}
	}

	return sb.String()
}

// TODO (tai): what magic is this
/*
func convert(s string, numRows int) string {
    if numRows <= 1 {return s}
    zigzag := make([][]byte, numRows)
    row, dir := 0,-1

    for i := 0; i < len(s); i++ {
        zigzag[row] = append(zigzag[row], s[i])
        if row == 0 || row == numRows -1 {
            dir *= -1
        }
        row += dir
    }

    return string( bytes.Join(zigzag,[]byte{} ))
}
*/
