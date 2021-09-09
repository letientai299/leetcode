package main

// All O`one Data Structure
//
// Hard
//
// https://leetcode.com/problems/all-oone-data-structure/
//
// Design a data structure to store the strings' count with the ability to
// return the strings with minimum and maximum counts.
//
// Implement the `AllOne` class:
//
// - `AllOne()` Initializes the object of the data structure.
// - `inc(String key)` Increments the count of the string `key` by `1`. If `key`
// does not exist in the data structure, insert it with count `1`.
// - `dec(String key)` Decrements the count of the string `key` by `1`. If the
// count of `key` is `0` after the decrement, remove it from the data structure.
// It is guaranteed that `key` exists in the data structure before the
// decrement.
// - `getMaxKey()` Returns one of the keys with the maximal count. If no element
// exists, return an empty string `""`.
// - `getMinKey()` Returns one of the keys with the minimum count. If no element
// exists, return an empty string `""`.
//
// **Example 1:**
//
// ```
// Input
// ["AllOne", "inc", "inc", "getMaxKey", "getMinKey", "inc", "getMaxKey",
// "getMinKey"]
// [[], ["hello"], ["hello"], [], [], ["leet"], [], []]
// Output
// [null, null, null, "hello", "hello", null, "hello", "leet"]
//
// Explanation
// AllOne allOne = new AllOne();
// allOne.inc("hello");
// allOne.inc("hello");
// allOne.getMaxKey(); // return "hello"
// allOne.getMinKey(); // return "hello"
// allOne.inc("leet");
// allOne.getMaxKey(); // return "hello"
// allOne.getMinKey(); // return "leet"
//
// ```
//
// **Constraints:**
//
// - `1 <= key.length <= 10`
// - `key` consists of lowercase English letters.
// - It is guaranteed that for each call to `dec`, `key` is existing in the data
// structure.
// - At most `5 * 104` calls will be made to `inc`, `dec`, `getMaxKey`, and
// `getMinKey`.

type AllOne struct {
	top, bot *AOLevel
	m        map[string]*AOLevel
}

type AOLevel struct {
	down, up *AOLevel
	count    int
	values   map[string]struct{}
}

func Constructor() AllOne {
	return AllOne{
		m: make(map[string]*AOLevel),
	}
}

func (a *AllOne) Inc(key string) {
	lvl, ok := a.m[key]

	if !ok {
		a.addNewKey(key)
		return
	}

	delete(lvl.values, key)

	if lvl.up == nil || lvl.up.count != lvl.count+1 {
		newLevel := &AOLevel{
			count:  lvl.count + 1,
			values: make(map[string]struct{}),
			down:   lvl,
			up:     lvl.up,
		}

		if lvl.up == nil {
			a.top = newLevel
		} else {
			lvl.up.down = newLevel
		}

		lvl.up = newLevel
	}

	a.m[key] = lvl.up
	lvl.up.values[key] = struct{}{}

	if len(lvl.values) != 0 {
		return
	}

	a.removeEmptyLevel(lvl)
}

func (a *AllOne) removeEmptyLevel(lvl *AOLevel) {
	if lvl.up == nil {
		a.top = lvl.down
	}

	if lvl.down == nil {
		a.bot = lvl.up
	}

	if a.top != nil {
		a.top.up = nil
	}

	if a.bot != nil {
		a.bot.down = nil
	}

	if lvl.down != nil {
		lvl.down.up = lvl.up
	}

	if lvl.up != nil {
		lvl.up.down = lvl.down
	}
}

func (a *AllOne) addNewKey(key string) {
	if a.bot != nil && a.bot.count == 1 {
		a.bot.values[key] = struct{}{}
		a.m[key] = a.bot
		return
	}

	// create a new level for keys with count of 1
	lvl := &AOLevel{
		count:  1,
		values: make(map[string]struct{}),
	}
	lvl.values[key] = struct{}{}
	a.m[key] = lvl

	if a.bot == nil {
		a.top = lvl
		a.bot = lvl
		return
	}

	a.bot.down = lvl
	lvl.up = a.bot
	a.bot = lvl
	return
}

func (a *AllOne) Dec(key string) {
	lvl, ok := a.m[key]
	if !ok {
		return
	}

	delete(lvl.values, key)

	if lvl.count == 1 {
		delete(a.m, key)

		if len(lvl.values) == 0 {
			a.removeEmptyLevel(lvl)
		}

		return
	}

	if lvl.down == nil || lvl.down.count != lvl.count-1 {
		newLevel := &AOLevel{
			count:  lvl.count - 1,
			values: make(map[string]struct{}),
			up:     lvl,
			down:   lvl.down,
		}

		if lvl.down == nil {
			a.bot = newLevel
		} else {
			lvl.down.up = newLevel
		}

		lvl.down = newLevel
	}

	lvl.down.values[key] = struct{}{}
	a.m[key] = lvl.down

	if len(lvl.values) == 0 {
		a.removeEmptyLevel(lvl)
	}
}

func (a *AllOne) GetMaxKey() string {
	if a.top == nil {
		return ""
	}
	for k := range a.top.values {
		return k
	}
	return ""
}

func (a *AllOne) GetMinKey() string {
	if a.bot == nil {
		return ""
	}

	for k := range a.bot.values {
		return k
	}
	return ""
}
