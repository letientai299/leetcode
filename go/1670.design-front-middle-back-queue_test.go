package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrontMiddleBackQueue(t *testing.T) {
	q := Constructor()

	n := q.head
	for n != nil {
		fmt.Print(n.val, " ")
		n = n.next
	}
	fmt.Println()

	// 4 3 2 1
	// 3 2 4 1
	/*

			["FrontMiddleBackQueue",
		"popMiddle",
		"pushMiddle",
		"pushMiddle",
		"popMiddle",
		"popMiddle",
		"popMiddle",
		"pushFront",
		"pushMiddle",
		"pushMiddle",
		"popMiddle",
		"pushBack",
		"popMiddle",
		"pushMiddle",
		"pushMiddle",
		"popMiddle",
		"pushMiddle",
		"popBack",
		"popMiddle",
		"popMiddle",
		"popMiddle",
		"pushMiddle",
		"popBack",
		"popMiddle",
		"popMiddle",
		"pushMiddle",
		"pushMiddle",
		"pushBack",
		"pushFront",
		"pushMiddle",
		"popMiddle"]

		[[],[],[207061],[41641],[],[],[],[944727],[188475],[933821],[],[411601],[],[804156],[991298],[],[363964],[],[],[],[],[536595],[],[],[],[950284],[902120],[883572],[388657],[485053],[]]

		[null,-1,null,null,41641,207061,-1,null,null,null,933821,null,944727,null,null,991298,null,411601,363964,933821,804156,null,536595,-1,-1,null,null,null,null,null,485053]
		[null,-1,null,null,41641,207061,-1,null,null,null,933821,null,944727,null,null,991298,null,411601,363964,188475,804156,null,536595,-1,-1,null,null,null,null,null,485053]

	*/
	q.PopMiddle()
	q.PushMiddle(207061)
	q.PushMiddle(41641)

	assert.Equal(t, q.PopMiddle(), 41641)
	assert.Equal(t, q.PopMiddle(), 207061)
	assert.Equal(t, q.PopMiddle(), -1)

	q.PushFront(944727)
	q.PushMiddle(188475)
	q.PushMiddle(933821)

	assert.Equal(t, q.PopMiddle(), 933821)

	q.PushBack(411601)
	assert.Equal(t, q.PopMiddle(), 944727)

	q.PushMiddle(804156)
	q.PushMiddle(991298)

	assert.Equal(t, q.PopMiddle(), 991298)
	q.PushMiddle(363964)
	assert.Equal(t, q.PopBack(), 411601)
	assert.Equal(t, q.PopMiddle(), 363964)
	assert.Equal(t, q.PopMiddle(), 188475)
	assert.Equal(t, q.PopMiddle(), 804156)

	q.PushMiddle(536595)
	assert.Equal(t, q.PopBack(), 536595)
	assert.Equal(t, q.PopMiddle(), -1)
	assert.Equal(t, q.PopMiddle(), -1)

	q.PushMiddle(950284)
	q.PushMiddle(902120)
	q.PushBack(883572)
	q.PushFront(388657)
	q.PushMiddle(485053)
	assert.Equal(t, q.PopMiddle(), 485053)
}
