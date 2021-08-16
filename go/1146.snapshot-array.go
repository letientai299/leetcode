package main

import "sort"

// Snapshot Array
//
// Medium
//
// https://leetcode.com/problems/snapshot-array/
//
// Implement a SnapshotArray that supports the following interface:
//
// - `SnapshotArray(int length)` initializes an array-like data structure with
// the given length.  **Initially, each element equals 0**.
// - `void set(index, val)` sets the element at the given `index` to be equal to
// `val`.
// - `int snap()` takes a snapshot of the array and returns the `snap_id`: the
// total number of times we called `snap()` minus `1`.
// - `int get(index, snap_id)` returns the value at the given `index`, at the
// time we took the snapshot with the given `snap_id`
//
// **Example 1:**
//
// ```
// Input: ["SnapshotArray","set","snap","set","get"]
// [[3],[0,5],[],[0,6],[0,0]]
// Output: [null,null,0,null,5]
// Explanation:
// SnapshotArray snapshotArr = new SnapshotArray(3); // set the length to be 3
// snapshotArr.set(0,5);  // Set array[0] = 5
// snapshotArr.snap();  // Take a snapshot, return snap_id = 0
// snapshotArr.set(0,6);
// snapshotArr.get(0,0);  // Get the value of array[0] with snap_id = 0, return
// 5
// ```
//
// **Constraints:**
//
// - `1 <= length <= 50000`
// - At most `50000` calls will be made to `set`, `snap`, and `get`.
// - `0 <= index < length`
// - `0 <= snap_id < `(the total number of times we call `snap()`)
// - `0 <= val <= 10^9`

type SnapshotArray struct {
	ver      int
	mem      []int
	values   [][]int
	versions [][]int
}

func Constructor1146(n int) SnapshotArray {
	return SnapshotArray{
		ver:      -1,
		mem:      make([]int, n),
		values:   make([][]int, n),
		versions: make([][]int, n),
	}
}

func (sa *SnapshotArray) Set(index int, val int) {
	sa.mem[index] = val
}

func (sa *SnapshotArray) Snap() int {
	sa.ver++
	for i, v := range sa.mem {
		l := len(sa.values[i])

		if l == 0 || sa.values[i][l-1] != v {
			sa.values[i] = append(sa.values[i], v)
			sa.versions[i] = append(sa.versions[i], sa.ver)
		} else {
			sa.versions[i][l-1] = sa.ver
		}
	}
	return sa.ver
}

func (sa *SnapshotArray) Get(index int, snapID int) int {
	arr := sa.versions[index]
	i := sort.SearchInts(arr, snapID)
	return sa.values[index][i]
}

// TODO (tai): TLE
