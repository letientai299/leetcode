package main

import "testing"

func TestCombinationIterator(t *testing.T) {
	ci := Constructor1286("abcdefg", 3)
	for i := 0; i < 35; i++ {
		t.Log(ci.HasNext(), ci.Next())
	}
	t.Log(ci.HasNext())
}
