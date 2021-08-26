package main

import "sort"

// Exam Room
//
// Medium
//
// https://leetcode.com/problems/exam-room/
//
// There is an exam room with `n` seats in a single row labeled from `0` to `n -
// 1`.
//
// When a student enters the room, they must sit in the seat that maximizes the
// distance to the closest person. If there are multiple such seats, they sit in
// the seat with the lowest number. If no one is in the room, then the student
// sits at seat number `0`.
//
// Design a class that simulates the mentioned exam room.
//
// Implement the `ExamRoom` class:
//
// - `ExamRoom(int n)` Initializes the object of the exam room with the number
// of the seats `n`.
// - `int seat()` Returns the label of the seat at which the next student will
// set.
// - `void leave(int p)` Indicates that the student sitting at seat `p` will
// leave the room. It is guaranteed that there will be a student sitting at seat
// `p`.
//
// **Example 1:**
//
// ```
// Input
// ["ExamRoom", "seat", "seat", "seat", "seat", "leave", "seat"]
// [[10], [], [], [], [], [4], []]
// Output
// [null, 0, 9, 4, 2, null, 5]
//
// Explanation
// ExamRoom examRoom = new ExamRoom(10);
// examRoom.seat(); // return 0, no one is in the room, then the student sits at
// seat number 0.
// examRoom.seat(); // return 9, the student sits at the last seat number 9.
// examRoom.seat(); // return 4, the student sits at the last seat number 4.
// examRoom.seat(); // return 2, the student sits at the last seat number 2.
// examRoom.leave(4);
// examRoom.seat(); // return 5, the student sits at the last seat number 5.
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 109`
// - It is guaranteed that there is a student sitting at seat `p`.
// - At most `104` calls will be made to `seat` and `leave`.

type ExamRoom struct {
	n      int
	stores []int
}

func Constructor(n int) ExamRoom {
	return ExamRoom{n: n}
}

func (ex *ExamRoom) Seat() int {
	switch len(ex.stores) {
	case 0:
		ex.stores = append(ex.stores, 0)
		return 0
	case 1:
		if ex.stores[0] == 0 {
			ex.stores = append(ex.stores, ex.n-1)
			return ex.n - 1
		}

		fallthrough
	default:
		return ex.find()
	}
}

func (ex *ExamRoom) Leave(p int) {
	i := sort.Search(len(ex.stores), func(i int) bool {
		return p <= ex.stores[i]
	})

	ex.stores = append(ex.stores[:i], ex.stores[i+1:]...)
}

func (ex *ExamRoom) find() int {
	bestDist := 0
	id := 0
	for i := 0; i < len(ex.stores)-1; i++ {
		a, b := ex.stores[i], ex.stores[i+1]
		if a+1 == b {
			continue
		}
		d := (b - a) / 2
		if d > bestDist {
			bestDist = d
			id = i
		}
	}

	if ex.stores[0] != 0 && ex.stores[0] >= bestDist {
		id = -1
		bestDist = ex.stores[0]
	}

	l := len(ex.stores)
	last := ex.stores[l-1]
	if last != ex.n-1 && (ex.n-1-last) > bestDist {
		id = l - 1
		bestDist = ex.n - 1 - last
	}

	switch id {
	case -1:
		ex.stores = append([]int{0}, ex.stores...)
		return 0
	case l - 1:
		ex.stores = append(ex.stores, ex.n-1)
		return ex.n - 1
	default:
		ex.stores = append(ex.stores, 0)
		copy(ex.stores[id+1:], ex.stores[id:])
		ex.stores[id+1] = ex.stores[id] + bestDist
		return ex.stores[id+1]
	}
}
