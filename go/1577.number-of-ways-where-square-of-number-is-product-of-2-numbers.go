package main

// medium
// https://leetcode.com/problems/number-of-ways-where-square-of-number-is-equal-to-product-of-two-numbers/

func numTriplets(nums1 []int, nums2 []int) int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)

	for _, v := range nums1 {
		m1[v]++
	}

	for _, v := range nums2 {
		m2[v]++
	}

	var r int
	for x, u := range m1 {
		x2 := x * x
		for y, v := range m2 {
			if x2%y == 0 {
				k := x2 / y
				if k == y {
					r += u * v * (m2[y] - 1)
				} else {
					r += u * v * (m2[k])
				}
			}

			y2 := y * y
			if y2%x != 0 {
				continue
			}
			k := y2 / x
			if k == x {
				r += u * v * (m1[x] - 1)
			} else {
				r += u * v * m1[k]
			}
		}
	}

	return r / 2
}
