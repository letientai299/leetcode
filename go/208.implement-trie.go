package main

// Implement Trie (Prefix Tree)
//
// Medium
//
// https://leetcode.com/problems/implement-trie-prefix-tree/
//
// A [**trie**](https://en.wikipedia.org/wiki/Trie) (pronounced as "try") or
// **prefix tree** is a tree data structure used to efficiently store and
// retrieve keys in a dataset of strings. There are various applications of this
// data structure, such as autocomplete and spellchecker.
//
// Implement the Trie class:
//
// - `Trie()` Initializes the trie object.
// - `void insert(String word)` Inserts the string `word` into the trie.
// - `boolean search(String word)` Returns `true` if the string `word` is in the
// trie (i.e., was inserted before), and `false` otherwise.
// - `boolean startsWith(String prefix)` Returns `true` if there is a previously
// inserted string `word` that has the prefix `prefix`, and `false` otherwise.
//
// **Example 1:**
//
// ```
// Input
// ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
// [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
// Output
// [null, null, true, false, true, null, true]
//
// Explanation
// Trie trie = new Trie();
// trie.insert("apple");
// trie.search("apple");   // return True
// trie.search("app");     // return False
// trie.startsWith("app"); // return True
// trie.insert("app");
// trie.search("app");     // return True
//
// ```
//
// **Constraints:**
//
// - `1 <= word.length, prefix.length <= 2000`
// - `word` and `prefix` consist only of lowercase English letters.
// - At most `3 * 104` calls **in total** will be made to `insert`, `search`,
// and `startsWith`.

// This implementation is slow because I hardly use any optimization related to
// Trie, also because I haven't really learn the data structure yet.
// TODO (tai): learn trie.

type Trie struct {
	kids [26]*Trie
	end  bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	if t == nil {
		return
	}

	if len(word) == 0 {
		t.end = true
		return
	}

	c := int(word[0] - 'a')
	if t.kids[c] == nil {
		kid := Constructor()
		t.kids[c] = &kid
	}

	kid := t.kids[c]
	kid.Insert(word[1:])
}

func (t *Trie) Search(word string) bool {
	if word == "" {
		return t.end
	}

	c := int(word[0] - 'a')
	kid := t.kids[c]
	return kid != nil && kid.Search(word[1:])
}

func (t *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return true
	}

	c := int(prefix[0] - 'a')
	kid := t.kids[c]
	return kid != nil && kid.StartsWith(prefix[1:])
}
