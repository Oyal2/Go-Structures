package AVLTree

import (
	"reflect"
	"testing"
)

func TestInsert1(t *testing.T) {
	list := New[int]()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	if actualValue := list.Length(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	if actualValue := list._root; actualValue.Value != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

}

func TestInsert2(t *testing.T) {
	list := New[int]()
	list.Insert(3)
	list.Insert(1)
	list.Insert(2)
	if actualValue := list.Length(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	if actualValue := list._root; actualValue.Value != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	arr := list.Traverse()
	if actualValue := arr[0]; actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue := arr[1]; actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue := arr[2]; actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

}

func TestInsert3(t *testing.T) {
	list := New[int]()
	list.Insert(3)
	list.Insert(2)
	list.Insert(1)
	if actualValue := list.Length(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	if actualValue := list._root; actualValue.Value != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

}

func TestInsert4(t *testing.T) {
	list := New[int]()
	list.Insert(1)
	list.Insert(3)
	list.Insert(2)
	if actualValue := list.Length(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	if actualValue := list._root; actualValue.Value != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
}
func TestRemove1(t *testing.T) {
	list := New[int]()
	list.Insert(1)
	list.Remove(1)
	if actualValue := list.Length(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestNew(t *testing.T) {
	tree := New[int]()

	if tree._root != nil {
		t.Errorf("Expected root to be nil, got %v", tree._root)
	}

	if tree._numberStored != 0 {
		t.Errorf("Expected number stored to be 0, got %d", tree._numberStored)
	}
}

func TestInsert(t *testing.T) {
	tree := New[int]()
	tree.Insert(5)

	if tree._root == nil {
		t.Errorf("Expected root to be non-nil, got %v", tree._root)
	}

	if tree._numberStored != 1 {
		t.Errorf("Expected number stored to be 1, got %d", tree._numberStored)
	}

	if tree._root.Value != 5 {
		t.Errorf("Expected root value to be 5, got %d", tree._root.Value)
	}
}

func TestRemove(t *testing.T) {
	tree := New[int]()
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(7)

	removed := tree.Remove(6)

	if removed != 6 {
		t.Errorf("Expected removed element to be 6, got %d", removed)
	}

	if tree._numberStored != 3 {
		t.Errorf("Expected number stored to be 3, got %d", tree._numberStored)
	}
}

func TestTraverse(t *testing.T) {
	tree := New[int]()
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(7)

	traversal := tree.Traverse()
	expectedTraversal := []int{4, 5, 6, 7}

	if !reflect.DeepEqual(traversal, expectedTraversal) {
		t.Errorf("Expected traversal to be %v, got %v", expectedTraversal, traversal)
	}
}

func TestContains(t *testing.T) {
	tree := New[int]()
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(7)

	if !tree.Contains(5) {
		t.Errorf("Expected tree to contain 5")
	}

	if tree.Contains(8) {
		t.Errorf("Expected tree to not contain 8")
	}
}

func TestLength(t *testing.T) {
	tree := New[int]()
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(7)

	if tree.Length() != 4 {
		t.Errorf("Expected length to be 4, got %d", tree.Length())
	}
}
