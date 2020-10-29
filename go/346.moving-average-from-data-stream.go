package main

/*
 * @lc app=leetcode id=346 lang=golang
 *
 * [346] Moving Average from Data Stream
 *
 * https://leetcode.com/problems/moving-average-from-data-stream/description/
 *
 * algorithms
 * Easy (68.19%)
 * Total Accepted:    151.2K
 * Total Submissions: 209.3K
 * Testcase Example:  '["MovingAverage","next","next","next","next"]\n[[3],[1],[10],[3],[5]]'
 *
 * Given a stream of integers and a window size, calculate the moving average
 * of all integers in the sliding window.
 *
 * Example:
 *
 *
 * MovingAverage m = new MovingAverage(3);
 * m.next(1) = 1
 * m.next(10) = (1 + 10) / 2
 * m.next(3) = (1 + 10 + 3) / 3
 * m.next(5) = (10 + 3 + 5) / 3
 */
type MovingAverage struct {
	win   []int
	size  int
	p     int
	count int
	sum   int
}

/** Initialize your data structure here. */
func MovingAverageConstructor(size int) MovingAverage {
	return MovingAverage{
		size: size,
		win:  make([]int, size),
		p:    0,
	}
}

func (mv *MovingAverage) Next(val int) float64 {
	mv.sum = mv.sum - mv.win[mv.p] + val
	mv.win[mv.p] = val
	mv.p = (mv.p + 1) % mv.size
	if mv.count < mv.size {
		mv.count++
	}
	return float64(mv.sum) / float64(mv.count)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
