package main

import "github.com/kirk91/leetcode/assert"

type Node struct {
	val    rune
	term   bool
	childs map[rune]*Node
}

type Trie struct {
	root  *Node
	depth int
	size  int
}

func NewTrie() *Trie {
	return &Trie{
		root: &Node{childs: make(map[rune]*Node)},
	}
}

func (t *Trie) Insert(word string) {
	rs := []rune(word)
	node := t.root
	for _, r := range rs {
		if _, ok := node.childs[r]; !ok {
			node.childs[r] = &Node{
				val:    r,
				childs: make(map[rune]*Node),
			}
		}
		node = node.childs[r]
	}
	node.term = true
}

func (t *Trie) Search(word string) bool {
	node := t.find(t.root, word)
	if node == nil || !node.term {
		return false
	}
	return true
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.find(t.root, prefix) != nil
}

func (t *Trie) find(node *Node, word string) *Node {
	if node == nil || word == "" {
		return nil
	}

	n := node
	rs := []rune(word)
	for _, r := range rs {
		if _, ok := n.childs[r]; !ok {
			return nil
		}
		n = n.childs[r]
	}
	return n
}

func main() {
	t := NewTrie()
	t.Insert("apple")
	t.Insert("ap")
	t.Insert("hello")
	t.Search("apple")
	assert.MustEqual(false, t.Search(""))
	assert.MustEqual(true, t.Search("apple"))
	assert.MustEqual(false, t.Search("appled"))
	assert.MustEqual(true, t.Search("ap"))
	assert.MustEqual(true, t.StartsWith("a"))
	assert.MustEqual(false, t.StartsWith(""))
	assert.MustEqual(false, t.StartsWith("b"))
}
