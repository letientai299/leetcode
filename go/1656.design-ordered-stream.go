package main

import (
	"sort"
)

type OrderedStream struct {
	ids []int
	ss  []string
	p   int
}

func Constructor_OrderedStream(n int) OrderedStream {
	return OrderedStream{
		p: 1,
	}
}

func (o *OrderedStream) Insert(id int, value string) []string {
	i := sort.SearchInts(o.ids, id)
	o.ids = append(o.ids[:i], append([]int{id}, o.ids[i:]...)...)
	o.ss = append(o.ss[:i], append([]string{value}, o.ss[i:]...)...)

	x := sort.SearchInts(o.ids, o.p)
	if x >= len(o.ids) || o.ids[x] != o.p {
		return nil
	}

	y := x
	for y+1 < len(o.ids) && (o.ids[y+1] == o.ids[y]+1) {
		y++
	}
	res := o.ss[x : y+1]
	o.p = o.ids[y] + 1

	return res
}
