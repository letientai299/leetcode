package main

import "sort"

// Time Based Key-Value Store
//
// Medium
//
// https://leetcode.com/problems/time-based-key-value-store/
//
// Design a time-based key-value data structure that can store multiple values
// for the same key at different time stamps and retrieve the key's value at a
// certain timestamp.
//
// Implement the `TimeMap` class:
//
// - `TimeMap()` Initializes the object of the data structure.
// - `void set(String key, String value, int timestamp)` Stores the key `key`
// with the value `value `at the given time `timestamp`.
// - `String get(String key, int timestamp)` Returns a value such that `set` was
// called previously, with `timestamp_prev <= timestamp`. If there are multiple
// such values, it returns the value associated with the largest
// `timestamp_prev`. If there are no values, it returns `""`.
//
// **Example 1:**
//
// ```
// Input
// ["TimeMap", "set", "get", "get", "set", "get", "get"]
// [[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo",
// 4], ["foo", 5]]
// Output
// [null, null, "bar", "bar", null, "bar2", "bar2"]
//
// Explanation
// TimeMap timeMap = new TimeMap();
// timeMap.set("foo", "bar", 1);  // store the key "foo" and value "bar" along
// with timestamp = 1.
// timeMap.get("foo", 1);         // return "bar"
// timeMap.get("foo", 3);         // return "bar", since there is no value
// corresponding to foo at timestamp 3 and timestamp 2, then the only value is
// at timestamp 1 is "bar".
// timeMap.set("foo", "bar2", 4); // store the key "foo" and value "ba2r" along
// with timestamp = 4.
// timeMap.get("foo", 4);         // return "bar2"
// timeMap.get("foo", 5);         // return "bar2"
//
// ```
//
// **Constraints:**
//
// - `1 <= key.length, value.length <= 100`
// - `key` and `value` consist of lowercase English letters and digits.
// - `1 <= timestamp <= 107`
// - All the timestamps `timestamp` of `set` are strictly increasing.
// - At most `2 * 105` calls will be made to `set` and `get`.

type TimeMap struct {
	stores map[string]*TimeMapVal
}

type TimeMapVal struct {
	values []string
	times  []int
}

func (v *TimeMapVal) insert(value string, time int) {
	v.times = append(v.times, time)
	v.values = append(v.values, value)
}

func (v *TimeMapVal) get(time int) string {
	i := sort.SearchInts(v.times, time)
	if i >= len(v.times) || v.times[i] > time {
		i--
	}
	if i < 0 {
		return ""
	}

	return v.values[i]
}

func Constructor981() TimeMap {
	return TimeMap{
		stores: make(map[string]*TimeMapVal, 100),
	}
}

func (kv *TimeMap) Set(key string, value string, timestamp int) {
	v, ok := kv.stores[key]
	if !ok {
		v = &TimeMapVal{
			times:  make([]int, 0, 100),
			values: make([]string, 0, 100),
		}
		kv.stores[key] = v
	}
	v.insert(value, timestamp)
}

func (kv *TimeMap) Get(key string, timestamp int) string {
	v, ok := kv.stores[key]
	if !ok {
		return ""
	}

	return v.get(timestamp)
}
