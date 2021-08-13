package main

// Design Add and Search Words Data Structure
//
// Medium
//
// https://leetcode.com/problems/design-add-and-search-words-data-structure/
//
// Design a data structure that supports adding new words and finding if a
// string matches any previously added string.
//
// Implement the `WordDictionary` class:
//
// - `WordDictionary()` Initializes the object.
// - `void addWord(word)` Adds `word` to the data structure, it can be matched
// later.
// - `bool search(word)` Returns `true` if there is any string in the data
// structure that matches `word` or `false` otherwise. `word` may contain dots
// `'.'` where dots can be matched with any letter.
//
// **Example:**
//
// ```
// Input
// ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
// [[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
// Output
// [null,null,null,null,false,true,true,true]
//
// Explanation
// WordDictionary wordDictionary = new WordDictionary();
// wordDictionary.addWord("bad");
// wordDictionary.addWord("dad");
// wordDictionary.addWord("mad");
// wordDictionary.search("pad"); // return False
// wordDictionary.search("bad"); // return True
// wordDictionary.search(".ad"); // return True
// wordDictionary.search("b.."); // return True
//
// ```
//
// **Constraints:**
//
// - `1 <= word.length <= 500`
// - `word` in `addWord` consists lower-case English letters.
// - `word` in `search` consist of  `'.'` or lower-case English letters.
// - At most `50000` calls will be made to `addWord` and `search`.

// TODO (tai): slow, yeah unoptimized trie again.

type WordDictionary struct {
	kids [26]*WordDictionary
	end  bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (t *WordDictionary) AddWord(word string) {
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
	kid.AddWord(word[1:])
}

func (t *WordDictionary) Search(word string) bool {
	if word == "" {
		return t.end
	}

	x := word[0]
	remain := word[1:]

	if x != '.' {
		c := int(x - 'a')
		kid := t.kids[c]
		return kid != nil && kid.Search(remain)
	}

	for _, k := range t.kids {
		if k != nil && k.Search(remain) {
			return true
		}
	}

	return false
}
