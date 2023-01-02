package main

type Trie struct {
	tree map[byte]*Trie
	leaf bool
}

func Constructor() Trie {
	return Trie{
		tree: make(map[byte]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.leaf = true
		return
	}
	tree := this.tree[word[0]]
	if tree == nil {
		tree = &Trie{
			tree: make(map[byte]*Trie),
		}
		this.tree[word[0]] = tree
	}
	tree.Insert(word[1:])
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.leaf
	}
	tree := this.tree[word[0]]
	if tree == nil {
		return false
	}
	return tree.Search(word[1:])
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	tree := this.tree[prefix[0]]
	if tree == nil {
		return false
	}
	return tree.StartsWith(prefix[1:])
}
