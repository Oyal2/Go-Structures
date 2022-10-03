package test_list

import (
	"github.com/oyal2/Go-Structures/list"
	"math/rand"
	"testing"
	"time"
)

func TestSinglyLink(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(100)
	checkArr := make([]int, size)
	newDl := list.NewSinglyLinkedList()

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
	newSl := list.NewSinglyLinkedList()

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
	newDl := list.NewSinglyLinkedList()

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
