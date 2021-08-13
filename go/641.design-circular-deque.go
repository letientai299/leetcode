package main

// Design Circular Deque
//
// Medium
//
// https://leetcode.com/problems/design-circular-deque/
//
// Design your implementation of the circular double-ended queue (deque).
//
// Implement the `MyCircularDeque` class:
//
// - `MyCircularDeque(int k)` Initializes the deque with a maximum size of `k`.
// - `boolean insertFront()` Adds an item at the front of Deque. Returns `true`
// if the operation is successful, or `false` otherwise.
// - `boolean insertLast()` Adds an item at the rear of Deque. Returns `true` if
// the operation is successful, or `false` otherwise.
// - `boolean deleteFront()` Deletes an item from the front of Deque. Returns
// `true` if the operation is successful, or `false` otherwise.
// - `boolean deleteLast()` Deletes an item from the rear of Deque. Returns
// `true` if the operation is successful, or `false` otherwise.
// - `int getFront()` Returns the front item from the Deque. Returns `-1` if the
// deque is empty.
// - `int getRear()` Returns the last item from Deque. Returns `-1` if the deque
// is empty.
// - `boolean isEmpty()` Returns `true` if the deque is empty, or `false`
// otherwise.
// - `boolean isFull()` Returns `true` if the deque is full, or `false`
// otherwise.
//
// **Example 1:**
//
// ```
// Input
// ["MyCircularDeque", "insertLast", "insertLast", "insertFront", "insertFront",
// "getRear", "isFull", "deleteLast", "insertFront", "getFront"]
// [[3], [1], [2], [3], [4], [], [], [], [4], []]
// Output
// [null, true, true, true, false, 2, true, true, true, 4]
//
// Explanation
// MyCircularDeque myCircularDeque = new MyCircularDeque(3);
// myCircularDeque.insertLast(1);  // return True
// myCircularDeque.insertLast(2);  // return True
// myCircularDeque.insertFront(3); // return True
// myCircularDeque.insertFront(4); // return False, the queue is full.
// myCircularDeque.getRear();      // return 2
// myCircularDeque.isFull();       // return True
// myCircularDeque.deleteLast();   // return True
// myCircularDeque.insertFront(4); // return True
// myCircularDeque.getFront();     // return 4
//
// ```
//
// **Constraints:**
//
// - `1 <= k <= 1000`
// - `0 <= value <= 1000`
// - At most `2000` calls will be made to `insertFront`, `insertLast`,
// `deleteFront`, `deleteLast`, `getFront`, `getRear`, `isEmpty`, `isFull`.

type MyCircularDeque struct {
	stores []int // cap is n+1
	start  int   // at the first element
	end    int   // after the last element
}

func Constructor641(k int) MyCircularDeque {
	return MyCircularDeque{
		stores: make([]int, k+1),
	}
}

func (cq *MyCircularDeque) InsertFront(value int) bool {
	if cq.IsFull() {
		return false
	}

	cq.start = (cq.start - 1 + cap(cq.stores)) % cap(cq.stores)
	cq.stores[cq.start] = value
	return true
}

func (cq *MyCircularDeque) InsertLast(value int) bool {
	if cq.IsFull() {
		return false
	}

	cq.stores[cq.end] = value
	cq.end = (cq.end + 1) % cap(cq.stores)
	return true
}

func (cq *MyCircularDeque) DeleteFront() bool {
	if cq.IsEmpty() {
		return false
	}

	cq.start = (cq.start + 1) % cap(cq.stores)
	return true
}

func (cq *MyCircularDeque) DeleteLast() bool {
	if cq.IsEmpty() {
		return false
	}

	cq.end = (cq.end - 1 + cap(cq.stores)) % cap(cq.stores)
	return true
}

func (cq *MyCircularDeque) GetFront() int {
	if cq.IsEmpty() {
		return -1
	}
	return cq.stores[cq.start]
}

func (cq *MyCircularDeque) GetRear() int {
	if cq.IsEmpty() {
		return -1
	}
	n := cap(cq.stores)
	id := (cq.end - 1 + n) % n
	return cq.stores[id]
}

func (cq *MyCircularDeque) IsEmpty() bool {
	return cq.start == cq.end
}

func (cq *MyCircularDeque) IsFull() bool {
	return (cq.end+1)%cap(cq.stores) == cq.start
}
