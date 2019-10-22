package main

import "math"

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 *
 * https://leetcode.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (27.55%)
 * Total Accepted:    527.5K Total Submissions: 1.9M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 *
 * Find the median of the two sorted arrays. The overall run time complexity
 * should be O(log (m+n)).
 *
 * You may assume nums1 and nums2 cannot be both empty.
 *
 * Example 1:
 *
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * The median is 2.0
 *
 *
 * Example 2:
 *
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * The median is (2 + 3)/2 = 2.5
 *
 *
 */

// I can't solve this myself.
// Here's a better explanation of the solution given on leetcode
// https://medium.com/@hazemu/finding-the-median-of-2-sorted-arrays-in-logarithmic-time-1d3f2ecbeb46
func findMedianSortedArrays(A []int, B []int) float64 {
	// find the length of both arrays
	m := len(A)
	n := len(B)
	if m < n {
		m, n = n, m
		A, B = B, A
	}

	// length of the expected left half of the merge.
	// this should contains the median, if m+n is odd, hence, we add 1
	k := (m + n + 1) / 2

	// now, we check the min and max values A can contribute to set of k.
	max := k
	if max > m {
		max = m
	}

	min := k - n // if B contributes all of its values
	if min < 0 {
		min = 0
	}

	// we will do binary search to look for the correct number of values
	// that A should contributes to the first half
	var med int
	var mid int

	for min <= max {
		mid = (min + max + 1) / 2
		ai := mid - 1

		// invalid index for A, thus, B must contribute all
		if ai < 0 || ai >= m {
			med = B[k-1]
			break
		}

		// x is the largest value that A can contributes,
		// due to the way we calculate min and max, this can't be overflow
		x := A[ai]

		// x2 is next value of x
		// used to check against y to see if we should advanced mid
		x2 := math.MinInt32
		if ai+1 < m {
			x2 = A[ai+1]
		}

		// y is the largest value contributed from B
		y := math.MinInt32

		// y2 is the next value
		// used to check against x to know how should we adjust mid
		y2 := math.MinInt32

		bi := k - mid - 1
		if bi >= 0 {
			y = B[bi]
		}

		if bi+1 < n {
			y2 = B[bi+1]
		}

		med = x
		if med < y {
			med = y
			// since med is possibly contributed by B,
			// need to check with x2.
			if med > x2 {
				// med larger than x2, means A should contribute more
				min = mid + 1
				continue
			}

			// found the med, stop
			break
		}

		if med > y2 {
			// B should contributes more values
			max = mid - 1
			continue
		}

		// found the med, stop
		break
	}

	if (m+n)%2 != 0 {
		return float64(med)
	}

	// if the merge has even numbers, need 1 more value to calculate the real median.
	var med2 int

	switch {
	case mid == 0 && k-mid >= n:
		// A didn't contribute to the first half, but B has no more
		// so, need to pick med2 from A
		med2 = A[0]

	case k-mid == 0 && mid >= m:
		// B didn't contribute to the first half, but B has no more
		// so, need to pick med2 from A
		med2 = B[0]

	case mid >= m:
		// if A can't contribute any more values,
		// pick next from B
		med2 = B[k-mid]

	case k-mid >= n:
		// if B can't contribute any more values,
		// pick next from A
		med2 = A[mid]

	default:
		// pick the smallest between next of x in A and next of y in B
		med2 = A[mid]
		y2 := B[k-mid]
		if y2 < med2 {
			med2 = y2
		}
	}

	return float64(med+med2) / 2
}
