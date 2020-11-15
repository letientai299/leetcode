package main

/*
 * @lc app=leetcode id=706 lang=golang
 *
 * [706] Design HashMap
 *
 * https://leetcode.com/problems/design-hashmap/description/
 *
 * algorithms
 * Easy (59.28%)
 * Total Accepted:    120.3K
 * Total Submissions: 194K
 * Testcase Example:  '["MyHashMap","put","put","get","get","put","get", "remove", "get"]\n' +
  '[[],[1,1],[2,2],[1],[3],[2,1],[2],[2],[2]]'
 *
 * Design a HashMap without using any built-in hash table libraries.
 *
 * To be specific, your design should include these functions:
 *
 *
 * put(key, value) : Insert a (key, value) pair into the HashMap. If the value
 * already exists in the HashMap, update the value.
 * get(key): Returns the value to which the specified key is mapped, or -1 if
 * this map contains no mapping for the key.
 * remove(key) : Remove the mapping for the value key if this map contains the
 * mapping for the key.
 *
 *
 *
 * Example:
 *
 *
 * MyHashMap hashMap = new MyHashMap();
 * hashMap.put(1, 1);
 * hashMap.put(2, 2);
 * hashMap.get(1);            // returns 1
 * hashMap.get(3);            // returns -1 (not found)
 * hashMap.put(2, 1);          // update the existing value
 * hashMap.get(2);            // returns 1
 * hashMap.remove(2);          // remove the mapping for 2
 * hashMap.get(2);            // returns -1 (not found)
 *
 *
 *
 * Note:
 *
 *
 * All keys and values will be in the range of [0, 1000000].
 * The number of operations will be in the range of [1, 10000].
 * Please do not use the built-in HashMap library.
 *
 *
*/
type MyHashMap struct {
	stores [][]keyVal
}

type keyVal struct {
	key int
	val int
}

/** Initialize your data structure here. */
func Constructor_MyHashMap() MyHashMap {
	stores := make([][]keyVal, 9)
	for i := range stores {
		stores[i] = make([]keyVal, 0, 3)
	}

	return MyHashMap{stores: stores}
}

func (h MyHashMap) hash(v int) int {
	return v * 31 / 13 % len(h.stores)
}

func (h *MyHashMap) Put(key int, value int) {
	i := h.hash(key)
	sub := h.stores[i]
	for i, v := range sub {
		if key == v.key {
			if v.val != value {
				sub[i] = keyVal{key, value}
			}
			return
		}
	}

	if len(sub) < len(h.stores)/4 {
		h.stores[i] = append(h.stores[i], keyVal{key: key, val: value})
		return
	}

	h.grow()
	h.Put(key, value)
}

func (h *MyHashMap) Get(key int) int {
	i := h.hash(key)
	sub := h.stores[i]
	for _, v := range sub {
		if key == v.key {
			return v.val
		}
	}
	return -1
}

func (h *MyHashMap) grow() {
	n := len(h.stores)
	stores := make([][]keyVal, n*2)
	for i := range stores {
		stores[i] = make([]keyVal, 0, n/3)
	}

	next := MyHashMap{stores: stores}
	for _, sub := range h.stores {
		for _, v := range sub {
			next.Put(v.key, v.val)
		}
	}

	h.stores = next.stores
}

func (h *MyHashMap) Remove(key int) {
	i := h.hash(key)
	sub := h.stores[i]
	for x, v := range sub {
		if key == v.key {
			h.stores[i] = append(sub[:x], sub[x+1:]...)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
