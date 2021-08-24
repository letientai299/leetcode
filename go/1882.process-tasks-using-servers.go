package main

import (
	"container/heap"
)

// Process Tasks Using Servers
//
// Medium
//
// https://leetcode.com/problems/process-tasks-using-servers/
//
// You are given two **0-indexed** integer arrays `servers` and `tasks` of
// lengths `n`​​​​​​ and `m`​​​​​​ respectively. `servers[i]` is the **weight**
// of the `i​​​​​​th`​​​​ server, and `tasks[j]` is the **time needed** to
// process the `j​​​​​​th`​​​​ task **in seconds**.
//
// Tasks are assigned to the servers using a **task queue**. Initially, all
// servers are free, and the queue is **empty**.
//
// At second `j`, the `jth` task is **inserted** into the queue (starting with
// the `0th` task being inserted at second `0`). As long as there are free
// servers and the queue is not empty, the task in the front of the queue will
// be assigned to a free server with the **smallest weight**, and in case of a
// tie, it is assigned to a free server with the **smallest index**.
//
// If there are no free servers and the queue is not empty, we wait until a
// server becomes free and immediately assign the next task. If multiple servers
// become free at the same time, then multiple tasks from the queue will be
// assigned **in order of insertion** following the weight and index priorities
// above.
//
// A server that is assigned task `j` at second `t` will be free again at second
// `t + tasks[j]`.
//
// Build an array `ans`​​​​ of length `m`, where `ans[j]` is the **index** of
// the server the `j​​​​​​th` task will be assigned to.
//
// Return _the array_ `ans`​​​​.
//
// **Example 1:**
//
// ```
// Input: servers = [3,3,2], tasks = [1,2,3,2,1,2]
// Output: [2,2,0,2,1,2]
// Explanation: Events in chronological order go as follows:
// - At second 0, task 0 is added and processed using server 2 until second 1.
// - At second 1, server 2 becomes free. Task 1 is added and processed using
// server 2 until second 3.
// - At second 2, task 2 is added and processed using server 0 until second 5.
// - At second 3, server 2 becomes free. Task 3 is added and processed using
// server 2 until second 5.
// - At second 4, task 4 is added and processed using server 1 until second 5.
// - At second 5, all servers become free. Task 5 is added and processed using
// server 2 until second 7.
// ```
//
// **Example 2:**
//
// ```
// Input: servers = [5,1,4,3,2], tasks = [2,1,2,4,5,2,1]
// Output: [1,4,1,4,1,3,2]
// Explanation: Events in chronological order go as follows:
// - At second 0, task 0 is added and processed using server 1 until second 2.
// - At second 1, task 1 is added and processed using server 4 until second 2.
// - At second 2, servers 1 and 4 become free. Task 2 is added and processed
// using server 1 until second 4.
// - At second 3, task 3 is added and processed using server 4 until second 7.
// - At second 4, server 1 becomes free. Task 4 is added and processed using
// server 1 until second 9.
// - At second 5, task 5 is added and processed using server 3 until second 7.
// - At second 6, task 6 is added and processed using server 2 until second 7.
//
// ```
//
// **Constraints:**
//
// - `servers.length == n`
// - `tasks.length == m`
// - `1 <= n, m <= 2 * 105`
// - `1 <= servers[i], tasks[j] <= 2 * 105`
func assignTasks(servers []int, tasks []int) []int {
	// TODO (tai): solve this without TLE

	queue := make([]server1882, len(servers))
	for id, w := range servers {
		queue[id] = server1882{id: id, weight: w}
	}

	h := serverQueue1882(queue)
	heap.Init(&h)

	for i, t := range tasks {
		for j := len(queue) - 1; j >= 0; j-- {
			s := queue[j]
			if s.busyUntil > 0 {
				queue[j].busyUntil--
			}
		}

		heap.Init(&h)
		s := heap.Pop(&h).(server1882)
		s.busyUntil += t
		heap.Push(&h, s)
		tasks[i] = s.id
	}

	return tasks
}

type server1882 struct {
	id        int
	weight    int
	busyUntil int
}

type serverQueue1882 []server1882

func (q serverQueue1882) Len() int { return len(q) }

func (q serverQueue1882) Less(i, j int) bool {
	s, a := q[i], q[j]
	if s.busyUntil != a.busyUntil {
		return s.busyUntil < a.busyUntil
	}

	if s.weight != a.weight {
		return s.weight < a.weight
	}

	return s.id < a.id
}

func (q serverQueue1882) Swap(i, j int) { q[i], q[j] = q[j], q[i] }

func (q *serverQueue1882) Push(x interface{}) {
	*q = append(*q, x.(server1882))
}

func (q *serverQueue1882) Pop() interface{} {
	n := q.Len() - 1
	v := (*q)[n]
	*q = (*q)[:n]
	return v
}
