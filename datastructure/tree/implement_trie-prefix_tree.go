package main

// https://www.geeksforgeeks.org/trie-insert-and-search/
type Trie struct {
	children    [26]*Trie
	isEndOfWord bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if this == nil {
		return
	}

	p := this
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if p.children[index] == nil {
			p.children[index] = new(Trie)
		}
		p = p.children[index]
	}
	// mark last node as leaf
	p.isEndOfWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if this == nil {
		return false
	}

	p := this
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if p.children[index] == nil {
			return false
		}
		p = p.children[index]
	}
	return p != nil && p.isEndOfWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if this == nil {
		return false
	}

	p := this
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		if p.children[index] == nil {
			return false
		}
		p = p.children[index]
	}
	return p != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
