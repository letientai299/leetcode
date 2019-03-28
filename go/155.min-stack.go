package main

/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 *
 * https://leetcode.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (35.96%)
 * Total Accepted:    276.3K
 * Total Submissions: 768.3K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n[[],[-2],[0],[-3],[],[],[],[]]'
 *
 *
 * Design a stack that supports push, pop, top, and retrieving the minimum
 * element in constant time.
 *
 *
 * push(x) -- Push element x onto stack.
 *
 *
 * pop() -- Removes the element on top of the stack.
 *
 *
 * top() -- Get the top element.
 *
 *
 * getMin() -- Retrieve the minimum element in the stack.
 *
 *
 *
 *
 * Example:
 *
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> Returns -3.
 * minStack.pop();
 * minStack.top();      --> Returns 0.
 * minStack.getMin();   --> Returns -2.
 *
 *
 */
type MinStack struct {
	data []int
	cap  int
	size int
	mins []int
}

/** initialize your data structure here. */
func MsConstructor() MinStack {
	return MinStack{
		data: make([]int, 10),
		cap:  10,
		size: 0,
		mins: nil,
	}
}

func (s *MinStack) Push(x int) {
	if s.size == s.cap {
		s.cap *= 2
		data := make([]int, s.cap)
		copy(data, s.data)
		s.data = data
	}

	s.data[s.size] = x
	s.size++

	if len(s.mins) == 0 {
		s.mins = []int{x}
	} else {
		if x <= s.mins[len(s.mins)-1] {
			s.mins = append(s.mins, x)
		}
	}
}

func (s *MinStack) Pop() {
	if s.size == 0 {
		return
	}

	x := s.data[s.size-1]
	s.data[s.size-1] = 0
	s.size--

	if x == s.mins[len(s.mins)-1] {
		s.mins = s.mins[:len(s.mins)-1]
	}
}

func (s *MinStack) Top() int {
	if s.size == 0 {
		return 0
	}

	return s.data[s.size-1]
}

func (s *MinStack) GetMin() int {
	return s.mins[len(s.mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
