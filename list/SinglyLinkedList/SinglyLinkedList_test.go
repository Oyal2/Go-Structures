package SinglyLinkedList

import (
	"math/rand"
	"testing"
	"time"

	"github.com/oyal2/Go-Structures/queue/Queue"
)

func TestSinglyLink(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(100)
	checkArr := make([]int, size)
	newDl := New[int]()

	for i := 0; i < size; i++ {
		randomNum := rand.Intn(1000)
		err := newDl.Insert(newDl.Length(), randomNum)
		if err != nil {
			t.Error(err)
		}
		checkArr[i] = randomNum
	}

	for i := range checkArr {
		val, err := newDl.Get(i)
		if err != nil {
			t.Error(err)
		}
		if val != checkArr[i] {
			t.Errorf("New DoublyLinked List was inserted incorrectly %d != %d", val, checkArr[i])
		}
	}
}

func TestUpdateSinglyLink(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(100)
	checkArr := make([]int, size)
	newSl := New[int]()

	for i := 0; i < size; i++ {
		randomNum := rand.Intn(1000)
		err := newSl.Insert(newSl.Length(), randomNum)
		if err != nil {
			t.Error(err)
		}
		checkArr[i] = randomNum
	}

	for i := range checkArr {
		err := newSl.Update(i, i)
		if err != nil {
			t.Error(err)
		}

		val, err := newSl.Get(i)
		if err != nil {
			t.Error(err)
		}
		if val != i {
			t.Errorf("New DoublyLinked List was inserted incorrectly %d != %d", val, i)
		}
	}
}

func TestFirstRemoveSinglyLink(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(100)
	checkArr := make([]int, size)
	newDl := New[int]()

	for i := 0; i < size; i++ {
		randomNum := rand.Intn(1000)
		err := newDl.Insert(newDl.Length(), randomNum)
		if err != nil {
			t.Error(err)
		}
		checkArr[i] = randomNum
	}

	for i := range checkArr {
		retVal, err := newDl.Remove(0)
		if err != nil {
			t.Error(err)
		}
		var val interface{}

		if len(checkArr)-1 != i {
			val, err = newDl.Get(0)
			if err != nil {
				t.Error(err)
			}
		}

		if retVal != checkArr[i] {
			t.Errorf("New DoublyLinked List was removed incorrectly %d != %d", val, checkArr[i])
		}

		if len(checkArr)-1 != i {
			if val != checkArr[i+1] {
				t.Errorf("New DoublyLinked List was removed incorrectly %d != %d", val, checkArr[i])
			}
		}

	}
}

func Test_RemoveFirst(t *testing.T) {
	newDl := Queue.New[int](New[int]())

	newDl.Enqueue(1)
	newDl.Enqueue(2)
	element, _ := newDl.Dequeue()
	if element != 1 {
		t.Error("Not 1")
	}
	element, _ = newDl.Dequeue()
	if element != 2 {
		t.Error("Not 1")
	}
}

func TestListAdd(t *testing.T) {
	list := New[string]()
	list.Add("a")
	list.Add("b", "c")
	if actualValue := list.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.Length(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || ok != nil {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
}

func TestListAppendAndPrepend(t *testing.T) {
	list := New[string]()
	list.Add("b")
	list.Prepend("a")
	list.Append("c")
	if actualValue := list.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.Length(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := list.Get(0); actualValue != "a" || ok != nil {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
	if actualValue, ok := list.Get(1); actualValue != "b" || ok != nil {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || ok != nil {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
}

func TestListRemove(t *testing.T) {
	list := New[string]()
	list.Add("a")
	list.Add("b", "c")
	list.Remove(2)
	if actualValue, ok := list.Get(2); actualValue == "a" || ok == nil {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	list.Remove(1)
	list.Remove(0)
	list.Remove(0) // no effect
	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Length(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestListGet(t *testing.T) {
	list := New[string]()
	list.Add("a")
	list.Add("b", "c")
	if actualValue, ok := list.Get(0); actualValue != "a" || ok != nil {
		t.Errorf("Got %v expected %v", actualValue, "a")
	}
	if actualValue, ok := list.Get(1); actualValue != "b" || ok != nil {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || ok != nil {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
	if actualValue, ok := list.Get(3); actualValue != "" || ok == nil {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	list.Remove(0)
	if actualValue, ok := list.Get(0); actualValue != "b" || ok != nil {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
}

func TestListClear(t *testing.T) {
	list := New[string]()
	list.Add("e", "f", "g", "a", "b", "c", "d")
	list.Clear()
	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Length(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestListIndexOf(t *testing.T) {
	list := New[string]()

	expectedIndex := -1
	if index := list.IndexOf("a"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}

	list.Add("a")
	list.Add("b", "c")

	expectedIndex = 0
	if index := list.IndexOf("a"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}

	expectedIndex = 1
	if index := list.IndexOf("b"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}

	expectedIndex = 2
	if index := list.IndexOf("c"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}
}

func TestListInsert(t *testing.T) {
	list := New[string]()
	list.Insert(0, "b")
	list.Insert(0, "c")
	list.Insert(0, "a")
	list.Insert(10, "x") // ignore
	if actualValue := list.Length(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	list.Insert(3, "d") // append
	if actualValue := list.Length(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}
}
