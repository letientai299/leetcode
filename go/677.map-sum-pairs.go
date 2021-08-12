package main

// Map Sum Pairs
//
// Medium
//
// https://leetcode.com/problems/map-sum-pairs/
//
// Design a map that allows you to do the following:
//
// - Maps a string key to a given value.
// - Returns the sum of the values that have a key with a prefix equal to a
// given string.
//
// Implement the `MapSum` class:
//
// - `MapSum()` Initializes the `MapSum` object.
// - `void insert(String key, int val)` Inserts the `key-val` pair into the map.
// If the `key` already existed, the original `key-value` pair will be
// overridden to the new one.
// - `int sum(string prefix)` Returns the sum of all the pairs' value whose
// `key` starts with the `prefix`.
//
// **Example 1:**
//
// ```
// Input
// ["MapSum", "insert", "sum", "insert", "sum"]
// [[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
// Output
// [null, null, 3, null, 5]
//
// Explanation
// MapSum mapSum = new MapSum();
// mapSum.insert("apple", 3);
// mapSum.sum("ap");           // return 3 (apple = 3)
// mapSum.insert("app", 2);
// mapSum.sum("ap");           // return 5 (apple + app = 3 + 2 = 5)
//
// ```
//
// **Constraints:**
//
// - `1 <= key.length, prefix.length <= 50`
// - `key` and `prefix` consist of only lowercase English letters.
// - `1 <= val <= 1000`
// - At most `50` calls will be made to `insert` and `sum`.

type MapSum struct {
	children  [26]*MapSum
	isWordEnd bool
	val       int
}

func Constructor() MapSum {
	return MapSum{}
}

func (t *MapSum) Insert(w string, val int) {
	if len(w) == 0 {
		t.isWordEnd = true
		t.val = val
		return
	}

	c := w[0] - 'a'
	if t.children[c] == nil {
		t.children[c] = &MapSum{}
	}

	child := t.children[c]
	child.Insert(w[1:], val)
}

func (t *MapSum) Sum(prefix string) int {
	if prefix == "" {
		return t.sum()
	}

	c, sub := prefix[0]-'a', prefix[1:]
	child := t.children[c]
	if child == nil {
		return 0
	}

	return child.Sum(sub)
}

func (t *MapSum) sum() int {
	if t == nil {
		return 0
	}

	s := t.val
	for _, c := range t.children {
		s += c.sum()
	}
	return s
}
