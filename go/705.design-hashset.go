package main

/*
 * @lc app=leetcode id=705 lang=golang
 *
 * [705] Design HashSet
 *
 * https://leetcode.com/problems/design-hashset/description/
 *
 * algorithms
 * Easy (58.64%)
 * Total Accepted:    105.5K
 * Total Submissions: 163.8K
 * Testcase Example:  '["MyHashSet","add","add","contains","contains","add","contains","remove","contains"]\n' +
  '[[],[1],[2],[1],[3],[2],[2],[2],[2]]'
 *
 * Design a HashSet without using any built-in hash table libraries.
 *
 * To be specific, your design should include these functions:
 *
 *
 * add(value): Insert a value into the HashSet. 
 * contains(value) : Return whether the value exists in the HashSet or not.
 * remove(value): Remove a value in the HashSet. If the value does not exist in
 * the HashSet, do nothing.
 *
 *
 *
 * Example:
 *
 *
 * MyHashSet hashSet = new MyHashSet();
 * hashSet.add(1);        
 * hashSet.add(2);        
 * hashSet.contains(1);    // returns true
 * hashSet.contains(3);    // returns false (not found)
 * hashSet.add(2);          
 * hashSet.contains(2);    // returns true
 * hashSet.remove(2);          
 * hashSet.contains(2);    // returns false (already removed)
 *
 *
 *
 * Note:
 *
 *
 * All values will be in the range of [0, 1000000].
 * The number of operations will be in the range of [1, 10000].
 * Please do not use the built-in HashSet library.
 *
 *
*/
type MyHashSet struct {
	stores [][]int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	stores := make([][]int, 16)
	for i := range stores {
		stores[i] = make([]int, 0, 4)
	}
	return MyHashSet{stores: stores}
}

func (h MyHashSet) hash(v int) int {
	return v * 31 / 13 % len(h.stores)
}

func (h *MyHashSet) Add(key int) {
	i := h.hash(key)
	sub := h.stores[i]
	for _, v := range sub {
		if key == v {
			return
		}
	}

	if len(sub) < len(h.stores)/4 {
		h.stores[i] = append(h.stores[i], key)
		return
	}

	h.grow()
	h.Add(key)
}

func (h *MyHashSet) Remove(key int) {
	i := h.hash(key)
	sub := h.stores[i]
	for x, v := range sub {
		if key == v {
			h.stores[i] = append(sub[:x], sub[x+1:]...)
		}
	}
}

/** Returns true if this set contains the specified element */
func (h *MyHashSet) Contains(key int) bool {
	i := h.hash(key)
	sub := h.stores[i]
	for _, v := range sub {
		if key == v {
			return true
		}
	}
	return false
}

func (h *MyHashSet) grow() {
	n := len(h.stores)
	stores := make([][]int, n*2)
	for i := range stores {
		stores[i] = make([]int, 0, n/4)
	}

	next := MyHashSet{stores: stores}
	for _, sub := range h.stores {
		for _, v := range sub {
			next.Add(v)
		}
	}

	h.stores = next.stores
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
