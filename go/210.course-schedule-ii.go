package main

// Course Schedule II
//
// Medium
//
// https://leetcode.com/problems/course-schedule-ii/
//
// There are a total of `numCourses` courses you have to take, labeled from `0`
// to `numCourses - 1`. You are given an array `prerequisites` where
// `prerequisites[i] = [ai, bi]` indicates that you **must** take course `bi`
// first if you want to take course `ai`.
//
// - For example, the pair `[0, 1]`, indicates that to take course `0` you have
// to first take course `1`.
//
// Return _the ordering of courses you should take to finish all courses_. If
// there are many valid answers, return **any** of them. If it is impossible to
// finish all courses, return **an empty array**.
//
// **Example 1:**
//
// ```
// Input: numCourses = 2, prerequisites = [[1,0]]
// Output: [0,1]
// Explanation: There are a total of 2 courses to take. To take course 1 you
// should have finished course 0. So the correct course order is [0,1].
//
// ```
//
// **Example 2:**
//
// ```
// Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
// Output: [0,2,1,3]
// Explanation: There are a total of 4 courses to take. To take course 3 you
// should have finished both courses 1 and 2. Both courses 1 and 2 should be
// taken after you finished course 0.
// So one correct course order is [0,1,2,3]. Another correct ordering is
// [0,2,1,3].
//
// ```
//
// **Example 3:**
//
// ```
// Input: numCourses = 1, prerequisites = []
// Output: [0]
//
// ```
//
// **Constraints:**
//
// - `1 <= numCourses <= 2000`
// - `0 <= prerequisites.length <= numCourses * (numCourses - 1)`
// - `prerequisites[i].length == 2`
// - `0 <= ai, bi < numCourses`
// - `ai != bi`
// - All the pairs `[ai, bi]` are **distinct**.
func findOrder(numCourses int, prerequisites [][]int) []int {
	depends := make(map[int][]int)
	child := make(map[int][]int)
	for _, p := range prerequisites {
		depends[p[0]] = append(depends[p[0]], p[1])
		child[p[1]] = append(child[p[1]], p[0])
	}

	res := make([]int, 0, numCourses)
	studied := make(map[int]struct{})

	for i := 0; i < numCourses; i++ {
		ds, ok := depends[i]
		if !ok || len(ds) == 0 {
			studied[i] = struct{}{}
			res = append(res, i)
			continue
		}
	}

	if len(res) == 0 {
		return res
	}

	start := 0
	for len(res) < numCourses {
		more := 0
		for i := start; i < len(res); i++ {
			for _, c := range child[res[i]] {
				if _, ok := studied[c]; ok {
					continue
				}

				can := true
				for _, d := range depends[c] {
					if _, ok := studied[d]; !ok {
						can = false
						break
					}
				}

				if can {
					studied[c] = struct{}{}
					res = append(res, c)
					more++
				}
			}
		}

		if more == 0 {
			return nil
		}
		start += len(res) - more
	}

	return res
}
