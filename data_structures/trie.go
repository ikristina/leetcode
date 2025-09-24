package datastructures

// Trie ...
type Trie struct {
	Root *TrieNode
}

// TrieNode ...
type TrieNode struct {
	Children map[rune]*TrieNode
	IsEnd    bool
}

// NewTrie ...
func NewTrie() *Trie {
	return &Trie{Root: &TrieNode{Children: make(map[rune]*TrieNode)}}
}

// Insert ...
func (t *Trie) Insert(word string) {
	node := t.Root
	for _, c := range word {
		if _, ok := node.Children[c]; !ok {
			node.Children[c] = &TrieNode{Children: make(map[rune]*TrieNode)}
		}
		node = node.Children[c]
	}
	node.IsEnd = true
}

// Search ...
func (t *Trie) Search(word string) bool {
	node := t.Root
	for _, c := range word {
		if _, ok := node.Children[c]; !ok {
			return false
		}
		node = node.Children[c]
	}
	return node.IsEnd
}

// StartsWith ...
func (t *Trie) StartsWith(prefix string) bool {
	node := t.Root
	for _, c := range prefix {
		if _, ok := node.Children[c]; !ok {
			return false
		}
		node = node.Children[c]
	}
	return true
}
