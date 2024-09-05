package tree

import "iter"

// Tree is a binary search tree node data structure
type Tree[E ~int] struct {
	value       E
	left, right *Tree[E]
}

// New creates a binary search tree node with given value
func New[E ~int](value E) *Tree[E] {
	return &Tree[E]{value: value}
}

// Insert adds given element to the tree
func (t *Tree[E]) Insert(v E) *Tree[E] {
	if t == nil {
		return New(v)
	}
	if v < t.value {
		t.left = t.left.Insert(v)
	} else {
		t.right = t.right.Insert(v)
	}
	return t
}

// All creates a function iterator over elements in the tree
func (t *Tree[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		t.push(yield)
	}
}

// push travers tree inorder and applies yield function to each element
func (t *Tree[E]) push(yield func(E) bool) bool {
	if t == nil {
		return true
	}
	return t.left.push(yield) && yield(t.value) && t.right.push(yield)
}
