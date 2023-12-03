package trie

// TrieNode represents each node in the trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// NewTrieNode creates a new Trie node with initialized children map
func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

// Trie represents the trie and has a reference to the root node
type Trie struct {
	root *TrieNode
}

// NewTrie creates a new Trie with an initialized root node
func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

// Insert adds a word to the trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = NewTrieNode()
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// Search returns true if the word is found in the trie
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
    node := t.root
    for _, ch := range prefix {
        if _, ok := node.children[ch]; !ok {
            return false
        }
		node = node.children[ch]
    }
    return true
}