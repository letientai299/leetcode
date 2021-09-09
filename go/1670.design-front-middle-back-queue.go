package main

// Design Front Middle Back Queue
//
// Medium
//
// https://leetcode.com/problems/design-front-middle-back-queue/
//
// Design a queue that supports `push` and `pop` operations in the front,
// middle, and back.
//
// Implement the `FrontMiddleBack` class:
//
// - `FrontMiddleBack()` Initializes the queue.
// - `void pushFront(int val)` Adds `val` to the **front** of the queue.
// - `void pushMiddle(int val)` Adds `val` to the **middle** of the queue.
// - `void pushBack(int val)` Adds `val` to the **back** of the queue.
// - `int popFront()` Removes the **front** element of the queue and returns it.
// If the queue is empty, return `-1`.
// - `int popMiddle()` Removes the **middle** element of the queue and returns
// it. If the queue is empty, return `-1`.
// - `int popBack()` Removes the **back** element of the queue and returns it.
// If the queue is empty, return `-1`.
//
// **Notice** that when there are **two** middle position choices, the operation
// is performed on the **frontmost** middle position choice. For example:
//
// - Pushing `6` into the middle of `[1, 2, 3, 4, 5]` results in `[1, 2, 6, 3,
// 4, 5]`.
// - Popping the middle from `[1, 2, 3, 4, 5, 6]` returns `3` and results in
// `[1, 2, 4, 5, 6]`.
//
// **Example 1:**
//
// ```
// Input:
// ["FrontMiddleBackQueue", "pushFront", "pushBack", "pushMiddle", "pushMiddle",
// "popFront", "popMiddle", "popMiddle", "popBack", "popFront"]
// [[], [1], [2], [3], [4], [], [], [], [], []]
// Output:
// [null, null, null, null, null, 1, 3, 4, 2, -1]
//
// Explanation:
// FrontMiddleBackQueue q = new FrontMiddleBackQueue();
// q.pushFront(1);   // [1]
// q.pushBack(2);    // [1, 2]
// q.pushMiddle(3);  // [1, 3, 2]
// q.pushMiddle(4);  // [1, 4, 3, 2]
// q.popFront();     // return 1 -> [4, 3, 2]
// q.popMiddle();    // return 3 -> [4, 2]
// q.popMiddle();    // return 4 -> [2]
// q.popBack();      // return 2 -> []
// q.popFront();     // return -1 -> [] (The queue is empty)
//
// ```
//
// **Constraints:**
//
// - `1 <= val <= 109`
// - At most `1000` calls will be made to `pushFront`, `pushMiddle`, `pushBack`,
// `popFront`, `popMiddle`, and `popBack`.

type FrontMiddleBackQueue struct {
	n               int
	head, mid, tail *FMBNode
}

type FMBNode struct {
	val  int
	prev *FMBNode
	next *FMBNode
}

func Constructor1670() FrontMiddleBackQueue { return FrontMiddleBackQueue{} }

func (q *FrontMiddleBackQueue) PushFront(val int) {
	node := &FMBNode{val: val}
	q.n++

	if q.n == 1 {
		q.head, q.mid, q.tail = node, node, node
		return
	}

	node.next = q.head
	q.head.prev = node
	q.head = node
	if q.n%2 == 0 {
		q.mid = q.mid.prev
	}
}

func (q *FrontMiddleBackQueue) PushMiddle(val int) {
	node := &FMBNode{val: val}
	q.n++

	if q.n == 1 {
		q.head, q.mid, q.tail = node, node, node
		return
	}

	if q.n%2 == 0 {
		node.next = q.mid
		node.prev = q.mid.prev
		if q.mid.prev != nil {
			q.mid.prev.next = node
		} else {
			q.head = node
		}
		q.mid.prev = node
		q.mid = node
	} else {
		node.prev = q.mid
		node.next = q.mid.next
		if q.mid.next != nil {
			q.mid.next.prev = node
		} else {
			q.tail = node
		}
		q.mid.next = node
		q.mid = node
	}
}

func (q *FrontMiddleBackQueue) PushBack(val int) {
	node := &FMBNode{val: val}
	q.n++

	if q.n == 1 {
		q.head, q.mid, q.tail = node, node, node
		return
	}

	q.tail.next = node
	node.prev = q.tail
	q.tail = node
	if q.n%2 == 1 {
		q.mid = q.mid.next
	}
}

func (q *FrontMiddleBackQueue) PopFront() int {
	if q.n == 0 {
		return -1
	}
	res := q.head.val

	q.n--
	if q.n == 0 {
		q.head, q.mid, q.tail = nil, nil, nil
		return res
	}

	q.head = q.head.next
	q.head.prev = nil

	if q.n%2 == 1 {
		q.mid = q.mid.next
	}

	return res
}

func (q *FrontMiddleBackQueue) PopMiddle() int {
	if q.n == 0 {
		return -1
	}

	q.n--
	res := q.mid.val
	if q.n == 0 {
		q.head, q.mid, q.tail = nil, nil, nil
		return res
	}

	if q.n%2 == 0 {
		mid := q.mid.prev
		mid.next = q.mid.next
		if mid.next != nil {
			mid.next.prev = mid
		}
		if mid.prev != nil {
			mid.prev.next = mid
		}
		q.mid = mid
		if q.n == 2 {
			q.head = mid
		}
	} else {
		mid := q.mid.next
		mid.prev = q.mid.prev
		q.mid = mid
		if mid.next != nil {
			mid.next.prev = mid
		}
		if mid.prev != nil {
			mid.prev.next = mid
		}
		if q.n == 1 {
			q.head, q.mid, q.tail = mid, mid, mid
		}
	}

	return res
}

func (q *FrontMiddleBackQueue) PopBack() int {
	if q.n == 0 {
		return -1
	}

	q.n--
	res := q.tail.val

	if q.n == 0 {
		q.head, q.mid, q.tail = nil, nil, nil
		return res
	}

	q.tail = q.tail.prev
	if q.tail != nil {
		q.tail.next = nil
	}

	if q.n%2 == 0 {
		q.mid = q.mid.prev
	}

	return res
}
