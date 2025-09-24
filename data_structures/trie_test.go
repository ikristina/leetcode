package datastructures

import "testing"

func TestTrieInsert(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("world")
	trie.Insert("hi")
	trie.Insert("hola")

	tests := []struct {
		input    string
		expected bool
	}{
		{"hello", true},
		{"world", true},
		{"hi", true},
		{"hola", true},
	}

	for _, test := range tests {
		if trie.Search(test.input) != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, trie.Search(test.input))
		}
	}
}
