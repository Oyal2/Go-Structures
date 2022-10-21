package test_list

import (
	"fmt"
	"github.com/oyal2/Go-Structures/list"
	"math/rand"
	"testing"
	"time"
)

func TestDoublyLink(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(100)
	checkArr := make([]int, size)
	newDl := list.NewDoublyLinkedList()

	for i := 0; i < size; i++ {
		//randomNum := rand.Intn(1000)
		err := newDl.Insert(0, i)
		if err != nil {
			t.Error(err)
		}
		checkArr[i] = i
	}

	for i, j := 0, len(checkArr)-1; i < j; i, j = i+1, j-1 {
		checkArr[i], checkArr[j] = checkArr[j], checkArr[i]
	}

	for i := range checkArr {
		val, err := newDl.Get(i)
		if err != nil {
			t.Error(err)
		}
		checkItem := checkArr[i]
		if val != checkItem {
			t.Errorf("New DoublyLinked List was inserted incorrectly %d != %d", val, checkArr[i])
		}
	}
}

func TestCustom(t *testing.T) {
	newDl := list.NewDoublyLinkedList()

	for i := 0; i < 4; i++ {
		//randomNum := rand.Intn(1000)
		err := newDl.Insert(i, i)
		if err != nil {
			t.Error(err)
		}
	}
	newDl.Insert(1, 4)
	fmt.Println("goo")
}

func TestUpdate(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(100)
	checkArr := make([]int, size)
	newDl := list.NewDoublyLinkedList()

	for i := 0; i < size; i++ {
		randomNum := rand.Intn(1000)
		err := newDl.Insert(newDl.Length(), randomNum)
		if err != nil {
			t.Error(err)
		}
		checkArr[i] = randomNum
	}

	for i := range checkArr {
		err := newDl.Update(i, i)
		if err != nil {
			t.Error(err)
		}

		val, err := newDl.Get(i)
		if err != nil {
			t.Error(err)
		}
		if val != i {
			t.Errorf("New DoublyLinked List was inserted incorrectly %d != %d", val, i)
		}
	}
}

func TestFirstRemove(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(100)
	checkArr := make([]int, size)
	newDl := list.NewDoublyLinkedList()

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
