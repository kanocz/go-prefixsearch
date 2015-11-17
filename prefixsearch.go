package prefixsearch

import (
	"strings"
	"unicode"
)

// SearchTree is struct to handle search tree
type SearchTree struct {
	root *node
}

type node struct {
	value    interface{}
	childnum uint
	childs   map[rune]*node
}

// New creates new search tree
func New() *SearchTree {
	return &SearchTree{
		root: &node{childs: map[rune]*node{}},
	}
}

// Add one leaf to tree
func (tree *SearchTree) Add(key string, value interface{}) {
	current := tree.root

	for _, sym := range strings.ToLower(key) {
		next, ok := current.childs[sym]
		if !ok {
			newone := &node{childs: map[rune]*node{}}
			current.childs[sym] = newone
			next = newone
		}
		current = next
	}

	current.value = value
}

// AutoComplete returns autocomplete suggestions for given prefix
func (tree *SearchTree) AutoComplete(prefix string) []interface{} {
	// walk thru prefix symbols
	current := tree.root
	for _, sym := range prefix {
		var ok bool
		current, ok = current.childs[unicode.ToLower(sym)]
		if !ok {
			return []interface{}{}
		}
	}

	// we have found, now very stupid tree walk :)
	var result []interface{}
	current.recurse(func(v interface{}) {
		if nil != v {
			result = append(result, v)
		}
	})
	return result
}

// Search searches for value of key
func (tree *SearchTree) Search(key string) interface{} {
	current := tree.root
	for _, sym := range key {
		var ok bool
		current, ok = current.childs[unicode.ToLower(sym)]
		if !ok {
			return nil
		}
	}
	return current.value
}

func (n *node) recurse(callback func(interface{})) {
	callback(n.value)
	for _, v := range n.childs {
		v.recurse(callback)
	}
}
