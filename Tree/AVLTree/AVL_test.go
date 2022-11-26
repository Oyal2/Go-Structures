package AVLTree

import "testing"

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

func TestInsert5(t *testing.T) {
	list := New[int]()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(66)
	list.Insert(5)

	if actualValue := list.Length(); actualValue != 6 {
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

func TestRemove2(t *testing.T) {
	list := New[int]()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Remove(3)
	if actualValue := list.Length(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}
