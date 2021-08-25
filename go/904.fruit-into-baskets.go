package main

// Fruit Into Baskets
//
// Medium
//
// https://leetcode.com/problems/fruit-into-baskets/
//
// You are visiting a farm that has a single row of fruit trees arranged from
// left to right. The trees are represented by an integer array `fruits` where
// `fruits[i]` is the **type** of fruit the `ith` tree produces.
//
// You want to collect as much fruit as possible. However, the owner has some
// strict rules that you must follow:
//
// - You only have **two** baskets, and each basket can only hold a **single
// type** of fruit. There is no limit on the amount of fruit each basket can
// hold.
// - Starting from any tree of your choice, you must pick **exactly one fruit**
// from **every** tree (including the start tree) while moving to the right. The
// picked fruits must fit in one of your baskets.
// - Once you reach a tree with fruit that cannot fit in your baskets, you must
// stop.
//
// Given the integer array `fruits`, return _the **maximum** number of fruits
// you can pick_.
//
// **Example 1:**
//
// ```
// Input: fruits = [1,2,1]
// Output: 3
// Explanation: We can pick from all 3 trees.
//
// ```
//
// **Example 2:**
//
// ```
// Input: fruits = [0,1,2,2]
// Output: 3
// Explanation: We can pick from trees [1,2,2].
// If we had started at the first tree, we would only pick from trees [0,1].
//
// ```
//
// **Example 3:**
//
// ```
// Input: fruits = [1,2,3,2,2]
// Output: 4
// Explanation: We can pick from trees [2,3,2,2].
// If we had started at the first tree, we would only pick from trees [1,2].
//
// ```
//
// **Example 4:**
//
// ```
// Input: fruits = [3,3,3,1,2,1,1,2,3,3,4]
// Output: 5
// Explanation: We can pick from trees [1,2,1,1,2].
//
// ```
//
// **Constraints:**
//
// - `1 <= fruits.length <= 105`
// - `0 <= fruits[i] < fruits.length`
func totalFruit(fruits []int) int {
	var types []int
	start := 0
	best := 1
	for i, v := range fruits {
		if len(types) == 0 || len(types) == 1 && fruits[types[0]] != v {
			types = append(types, i)
			continue
		}

		switch {
		case v == fruits[types[0]]:
			types[0] = i
		case v == fruits[types[1]]:
			types[1] = i
		default:
			// new type
			if i-start > best {
				best = i - start
			}

			if types[0] > types[1] {
				start = types[1] + 1
				types[1] = i
			} else {
				start = types[0] + 1
				types[0] = i
			}
		}
	}

	if len(fruits)-start > best {
		best = len(fruits) - start
	}
	return best
}
