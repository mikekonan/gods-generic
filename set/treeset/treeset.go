package treeset

import (
	"fmt"
	"strings"

	"github.com/mikekonan/gods-generic/tree/redblacktree"
	"github.com/mikekonan/gods-generic/utils"
)

// Set holds elements in a red-black tree
type Set[V any] struct {
	tree *redblacktree.Tree[V, struct{}]
}

var itemExists = struct{}{}

// NewWithComparator instantiates a new empty set with the custom comparator.
func NewWithComparator[V any](comparator utils.Comparator[V]) *Set[V] {
	set := &Set[V]{tree: redblacktree.NewWithComparator[V, struct{}](comparator)}

	return set
}

// Add adds the items (one or more) to the set.
func (set *Set[V]) Add(items ...V) {
	for _, item := range items {
		set.tree.Put(item, itemExists)
	}
}

// AddIfFunc adds the items (one or more) to the set based on ifFunc.
func (set *Set[V]) AddIfFunc(item V, ifFunc func(V, V) bool) {
	set.tree.PutIfFunc(item, itemExists, ifFunc)
}

// Remove removes the items (one or more) from the set.
func (set *Set[V]) Remove(items ...V) {
	for _, item := range items {
		set.tree.Remove(item)
	}
}

// Contains checks weather items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set[V]) Contains(items ...V) bool {
	for _, item := range items {
		if _, contains := set.tree.Get(item); !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Set[V]) Empty() bool {
	return set.tree.Size() == 0
}

// Size returns number of elements within the set.
func (set *Set[V]) Size() int {
	return set.tree.Size()
}

// Clear clears all values in the set.
func (set *Set[V]) Clear() {
	set.tree.Clear()
}

// Values returns all items in the set.
func (set *Set[V]) Values() []V {
	return set.tree.Keys()
}

// ReversedValues returns all items in the set in reversed order.
func (set *Set[V]) ReversedValues() []V {
	return set.tree.ReversedKeys()
}

// String returns a string representation of container
func (set *Set[V]) String() string {
	str := "TreeSet\n"
	var items []string
	for _, v := range set.tree.Keys() {
		items = append(items, fmt.Sprintf("%v", v))
	}
	str += strings.Join(items, ", ")
	return str
}

// First returns the minimum key and its value from the tree map.
// Returns nil, nil if map is empty.
func (m *Set[V]) First() (value V) {
	if node := m.tree.Left(); node != nil {
		return node.Key
	}

	return
}

// Last returns the minimum key and its value from the tree map.
// Returns nil, nil if map is empty.
func (m *Set[V]) Last() (value V) {
	if node := m.tree.Right(); node != nil {
		return node.Key
	}

	return
}
