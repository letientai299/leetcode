package main

// medium
// 1381. Design a Stack With Increment Operation

type CustomStack struct {
	store []int
	incs  []int
	i     int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		store: make([]int, maxSize),
		incs:  make([]int, maxSize),
	}
}

func (c *CustomStack) Push(x int) {
	if c.i < len(c.store) {
		c.store[c.i] = x
		c.i++
	}
}

func (c *CustomStack) Pop() int {
	if c.i == 0 {
		return -1
	}

	c.i--
	p := c.incs[c.i]
	if c.i > 0 {
		c.incs[c.i-1] += p
	}
	c.incs[c.i] = 0
	return c.store[c.i] + p
}

func (c *CustomStack) Increment(k int, val int) {
	if c.i == 0 {
		return
	}

	if k > c.i {
		c.incs[c.i-1] += val
	} else {
		c.incs[k-1] += val
	}
}
