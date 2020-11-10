package main

import "sort"

/*
 * @lc app=leetcode id=1086 lang=golang
 *
 * [1086] High Five
 *
 * https://leetcode.com/problems/high-five/description/
 *
 * algorithms
 * Easy (75.58%)
 * Total Accepted:    40.6K
 * Total Submissions: 50.7K
 * Testcase Example:  '[[1,91],[1,92],[2,93],[2,97],[1,60],[2,77],[1,65],[1,87],[1,100],[2,100],[2,76]]'
 *
 * Given a list of the scores of different students, items, where items[i] =
 * [IDi, scorei] represents one score from a student with IDi, calculate each
 * student's top five average.
 *
 * Return the answer as an array of pairs result, where result[j] = [IDj,
 * topFiveAveragej] represents the student with IDj and their top five average.
 * Sort result by IDj in increasing order.
 *
 * A student's top five average is calculated by taking the sum of their top
 * five scores and dividing it by 5 using integer division.
 *
 *
 * Example 1:
 *
 *
 * Input: items =
 * [[1,91],[1,92],[2,93],[2,97],[1,60],[2,77],[1,65],[1,87],[1,100],[2,100],[2,76]]
 * Output: [[1,87],[2,88]]
 * Explanation:
 * The student with ID = 1 got scores 91, 92, 60, 65, 87, and 100. Their top
 * five average is (100 + 92 + 91 + 87 + 65) / 5 = 87.
 * The student with ID = 2 got scores 93, 97, 77, 100, and 76. Their top five
 * average is (100 + 97 + 93 + 77 + 76) / 5 = 88.6, but with integer division
 * their average converts to 88.
 *
 *
 * Example 2:
 *
 *
 * Input: items =
 * [[1,100],[7,100],[1,100],[7,100],[1,100],[7,100],[1,100],[7,100],[1,100],[7,100]]
 * Output: [[1,100],[7,100]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= items.length <= 1000
 * items[i].length == 2
 * 1 <= IDi <= 1000
 * 0 <= scorei <= 100
 * For each IDi, there will be at least five scores.
 *
 *
 */
func highFive(items [][]int) [][]int {
	m := make(map[int][]int)
	var ids []int
	for _, item := range items {
		id, v := item[0], item[1]
		if _, ok := m[id]; !ok {
			m[id] = make([]int, 5)
			ids = append(ids, id)
		}

		// m[id] is decreasing top 5 of student ID.
		loc := -1
		for i, x := range m[id] {
			if x < v {
				loc = i
				break
			}
		}
		if loc < 0 {
			continue
		}

		for i := loc; i < 5; i++ {
			m[id][i], v = v, m[id][i]
		}
	}

	res := make([][]int, 0, len(ids))
	sort.Ints(ids)
	for _, id := range ids {
		avg := 0
		for _, x := range m[id] {
			avg += x
		}
		avg /= 5
		res = append(res, []int{id, avg})
	}
	return res
}
