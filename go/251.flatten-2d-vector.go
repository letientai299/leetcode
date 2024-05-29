package main

// Flatten 2D Vector
//
// # Medium
//
// https://leetcode.com/problems/flatten-2d-vector
type Vector2D struct {
	data [][]int
	r    int
	c    int
}

func Constructor(vec [][]int) Vector2D {
	end := 0
	for i, v := range vec {
		if len(vec[i]) != 0 {
			vec[end] = v
			end++
		}
	}

	v := Vector2D{
		data: vec[:end],
		r:    0,
		c:    0,
	}

	return v
}

func (v *Vector2D) Next() int {
	d := v.data[v.r][v.c]
	v.c++
	if v.c >= len(v.data[v.r]) {
		v.r++
		v.c = 0
	}
	return d
}

func (v *Vector2D) HasNext() bool {
	return v.r < len(v.data) && v.c < len(v.data[v.r])
}
