package main

/*
 * @lc app=leetcode id=284 lang=golang
 *
 * [284] Peeking Iterator
 *
 * https://leetcode.com/problems/peeking-iterator/description/
 *
 * algorithms
 * Medium (42.50%)
 * Total Accepted:    109.7K
 * Total Submissions: 232.2K
 * Testcase Example:  '["PeekingIterator","next","peek","next","next","hasNext"]\n' +
  '[[[1,2,3]],[],[],[],[],[]]'
 *
 * Given an Iterator class interface with methods: next() and hasNext(), design
 * and implement a PeekingIterator that support the peek() operation -- it
 * essentially peek() at the element that will be returned by the next call to
 * next().
 *
 * Example:
 *
 *
 * Assume that the iterator is initialized to the beginning of the list:
 * [1,2,3].
 *
 * Call next() gets you 1, the first element in the list.
 * Now you call peek() and it returns 2, the next element. Calling next() after
 * that still return 2.
 * You call next() the final time and it returns 3, the last element.
 * Calling hasNext() after that should return false.
 *
 *
 * Follow up: How would you extend your design to be generic and work with all
 * types, not just integer?
 *
*/
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type Iterator interface {
	hasNext() bool
	next() int
}

type PeekingIterator struct {
	iter     Iterator
	vNext    int
	vHasNext bool
}

func Constructor284(iter Iterator) *PeekingIterator {
	pi := PeekingIterator{
		iter:     iter,
		vHasNext: iter.hasNext(),
	}
	if pi.vHasNext {
		pi.vNext = iter.next()
	}
	return &pi
}

func (pi *PeekingIterator) hasNext() bool {
	return pi.vHasNext
}

func (pi *PeekingIterator) next() int {
	v := pi.vNext
	pi.vHasNext = pi.iter.hasNext()
	if pi.vHasNext {
		pi.vNext = pi.iter.next()
	}

	return v
}

func (pi *PeekingIterator) peek() int {
	return pi.vNext
}
