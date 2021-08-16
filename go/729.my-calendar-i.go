package main

// My Calendar I
//
// Medium
//
// https://leetcode.com/problems/my-calendar-i/
//
// You are implementing a program to use as your calendar. We can add a new
// event if adding the event will not cause a **double booking**.
//
// A **double booking** happens when two events have some non-empty intersection
// (i.e., some moment is common to both events.).
//
// The event can be represented as a pair of integers `start` and `end` that
// represents a booking on the half-open interval `[start, end)`, the range of
// real numbers `x` such that `start <= x < end`.
//
// Implement the `MyCalendar` class:
//
// - `MyCalendar()` Initializes the calendar object.
// - `boolean book(int start, int end)` Returns `true` if the event can be added
// to the calendar successfully without causing a **double booking**. Otherwise,
// return `false` and do not add the event to the calendar.
//
// **Example 1:**
//
// ```
// Input
// ["MyCalendar", "book", "book", "book"]
// [[], [10, 20], [15, 25], [20, 30]]
// Output
// [null, true, false, true]
//
// Explanation
// MyCalendar myCalendar = new MyCalendar();
// myCalendar.book(10, 20); // return True
// myCalendar.book(15, 25); // return False, It can not be booked because time
// 15 is already booked by another event.
// myCalendar.book(20, 30); // return True, The event can be booked, as the
// first event takes every time less than 20, but not including 20.
// ```
//
// **Constraints:**
//
// - `0 <= start < end <= 109`
// - At most `1000` calls will be made to `book`.

type MyCalendar struct {
	head *MyCalendarEvent
}

type MyCalendarEvent struct {
	start, end int
	next       *MyCalendarEvent
}

func Constructor729() MyCalendar {
	return MyCalendar{}
}

func (ca *MyCalendar) Book(start int, end int) bool {
	if ca.head == nil {
		ca.head = &MyCalendarEvent{
			start: start,
			end:   end,
		}
		return true
	}

	if ca.head.start == end {
		ca.head.start = start
		return true
	}

	if ca.head.start > end {
		ca.head = &MyCalendarEvent{
			start: start,
			end:   end,
			next:  ca.head,
		}
		return true
	}

	pre := ca.head
	n := ca.head.next
	for n != nil {
		if n.start < start {
			pre = n
			n = pre.next
			continue
		}

		if pre.end > start || n.start < end {
			return false
		}

		if pre.end == start && n.start == end {
			pre.end = n.end
			pre.next = n.next
			return true
		}

		if pre.end == start {
			pre.end = end
			return true
		}

		if n.start == end {
			n.start = start
			return true
		}

		pre.next = &MyCalendarEvent{start: start, end: end, next: n}
		return true
	}

	if pre.end > start {
		return false
	}

	pre.next = &MyCalendarEvent{start: start, end: end}
	return true
}
