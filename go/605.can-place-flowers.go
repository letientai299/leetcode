package main

/*
 * @lc app=leetcode id=605 lang=golang
 *
 * [605] Can Place Flowers
 *
 * https://leetcode.com/problems/can-place-flowers/description/
 *
 * algorithms
 * Easy (30.98%)
 * Total Accepted:    63.9K
 * Total Submissions: 206.1K
 * Testcase Example:  '[1,0,0,0,1]\n1'
 *
 * Suppose you have a long flowerbed in which some of the plots are planted and
 * some are not. However, flowers cannot be planted in adjacent plots - they
 * would compete for water and both would die.
 *
 * Given a flowerbed (represented as an array containing 0 and 1, where 0 means
 * empty and 1 means not empty), and a number n, return if n new flowers can be
 * planted in it without violating the no-adjacent-flowers rule.
 *
 * Example 1:
 *
 * Input: flowerbed = [1,0,0,0,1], n = 1
 * Output: True
 *
 *
 *
 * Example 2:
 *
 * Input: flowerbed = [1,0,0,0,1], n = 2
 * Output: False
 *
 *
 *
 * Note:
 *
 * The input array won't violate no-adjacent-flowers rule.
 * The input array size is in the range of [1, 20000].
 * n is a non-negative integer which won't exceed the input array size.
 *
 *
 */
func canPlaceFlowers(flowerbed []int, n int) bool {
	slots := 0

	size := len(flowerbed)
	canPlant := func(i int) bool {
		if flowerbed[i] == 1 {
			return false
		}

		if i-1 >= 0 && flowerbed[i-1] == 1 {
			return false
		}

		if i+1 < size && flowerbed[i+1] == 1 {
			return false
		}

		return true
	}

	for i := 0; i < size; i++ {
		if canPlant(i) {
			slots++
			flowerbed[i] = 1
		}
	}

	return n <= slots
}
