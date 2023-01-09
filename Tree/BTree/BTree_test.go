package BTree

import (
	"testing"
)

func TestBTree(t *testing.T) {
	// Create a new B-Tree with a maximum of 3 children per node
	tree := New(func(a, b int) bool {
		return a < b
	}, 3)

	// Insert some elements into the tree
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)

	// Check that the number of elements stored in the tree is correct
	if tree._numberStored != 5 {
		t.Errorf("Expected number of elements to be 5, got %d", tree._numberStored)
	}

	// Check the structure of the tree
	if tree._root.N != 2 {
		t.Errorf("Expected root node to have 2 keys, got %d", tree._root.N)
	}
	if tree._root.Leaf != false {
		t.Errorf("Expected root node to be an internal node, got leaf node")
	}
	if tree._root.Children[0].N != 1 {
		t.Errorf("Expected root's first child to have 1 key, got %d", tree._root.Children[0].N)
	}
	if tree._root.Children[1].N != 1 {
		t.Errorf("Expected root's second child to have 1 key, got %d", tree._root.Children[1].N)
	}
	if tree._root.Children[2].N != 1 {
		t.Errorf("Expected root's third child to have 1 key, got %d", tree._root.Children[2].N)
	}
}
