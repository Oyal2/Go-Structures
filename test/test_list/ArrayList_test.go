package test_list

import (
	"Go-Structures/list"
	"math/rand"
	"time"

	"testing"
)

//Test the ArrayList when its fully filled
func TestFillArray(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(100)
	checkArr := make([]int, size)
	newArr := list.NewArrayList(size)

	for i := 0; i < size; i++ {
		randomNum := rand.Intn(1000)
		err := newArr.Insert(newArr.Length(), randomNum)
		if err != nil {
			t.Error(err)
		}
		checkArr[i] = randomNum
	}

	for i := range checkArr {
		val, err := newArr.Get(i)
		if err != nil {
			t.Error(err)
		}
		if val != checkArr[i] {
			t.Errorf("New ArrayList was inserted incorrectly %d != %d", val, checkArr[i])
		}
	}
}

//Test the get function of the ArrayList
func TestGetIndex(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(100)
	checkArr := make([]int, size+1)
	newArr := list.NewArrayList(size)

	for i := 0; i < size+1; i++ {
		randomNum := rand.Intn(1000)
		err := newArr.Insert(newArr.Length(), randomNum)
		if err != nil {
			t.Error(err)
		}
		checkArr[i] = randomNum
	}

	for i := range checkArr {
		val, err := newArr.Get(i)
		if err != nil {
			t.Error(err)
		}
		if val != checkArr[i] {
			t.Errorf("New ArrayList was inserted incorrectly %d != %d", val, checkArr[i])
		}
	}
}

//Test the ArrayList when inserted more than initialized
func TestOverFilledArray(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(100)
	checkArr := make([]int, size+1)
	newArr := list.NewArrayList(size)

	for i := 0; i < size+1; i++ {
		randomNum := rand.Intn(1000)
		err := newArr.Insert(newArr.Length(), randomNum)
		if err != nil {
			t.Error(err)
		}
		checkArr[i] = randomNum
	}

	for i := range checkArr {
		val, err := newArr.Get(i)
		if err != nil {
			t.Error(err)
		}
		if val != checkArr[i] {
			t.Errorf("New ArrayList was inserted incorrectly %d != %d", val, checkArr[i])
		}
	}
}

//Test a filled ArrayList and delete the middle index
func TestDeleteMiddleIndexArray(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(100)
	checkArr := make([]int, size)
	newArr := list.NewArrayList(size)

	for i := 0; i < size; i++ {
		err := newArr.Insert(newArr.Length(), i)
		if err != nil {
			t.Error(err)
		}
		checkArr[i] = i
	}

	removedVal, err := newArr.Remove(size / 2)
	if err != nil {
		t.Error(err)
	}
	if removedVal != checkArr[size/2] {
		t.Errorf("New ArrayList did not delete the right number: %d", removedVal)
	}

	if newArr.Length() != size-1 {
		t.Errorf("New ArrayList length: %d", newArr.Length())
	}

	for i := 0; i < newArr.Length(); i++ {
		val, err := newArr.Get(i)
		if err != nil {
			t.Error(err)
		}
		if val == removedVal {
			t.Errorf("New ArrayList did not delete %d", val)
		}
	}
}

//Test a filled ArrayList and delete the first index
func TestDeleteFirstIndexArray(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(100)
	checkArr := make([]int, size)
	newArr := list.NewArrayList(size)

	for i := 0; i < size; i++ {
		err := newArr.Insert(newArr.Length(), i)
		if err != nil {
			t.Error(err)
		}
		checkArr[i] = i
	}

	removedVal, err := newArr.Remove(0)
	if err != nil {
		t.Error(err)
	}
	if removedVal != checkArr[0] {
		t.Errorf("New ArrayList did not delete the right number: %d", removedVal)
	}

	if newArr.Length() != size-1 {
		t.Errorf("New ArrayList length: %d", newArr.Length())
	}

	for i := 0; i < newArr.Length(); i++ {
		val, err := newArr.Get(i)
		if err != nil {
			t.Error(err)
		}
		if val == removedVal {
			t.Errorf("New ArrayList did not delete %d", val)
		}
	}
}

//Test a filled ArrayList and delete the last index
func TestDeleteLastIndexArray(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(100)
	checkArr := make([]int, size)
	newArr := list.NewArrayList(size)

	for i := 0; i < size; i++ {
		err := newArr.Insert(newArr.Length(), i)
		if err != nil {
			t.Error(err)
		}
		checkArr[i] = i
	}

	removedVal, err := newArr.Remove(size - 1)
	if err != nil {
		t.Error(err)
	}
	if removedVal != checkArr[size-1] {
		t.Errorf("New ArrayList did not delete the right number: %d", removedVal)
	}

	if newArr.Length() != size-1 {
		t.Errorf("New ArrayList length: %d", newArr.Length())
	}

	for i := 0; i < newArr.Length(); i++ {
		val, err := newArr.Get(i)
		if err != nil {
			t.Error(err)
		}
		if val == removedVal {
			t.Errorf("New ArrayList did not delete %d", val)
		}
	}
}

//Test to update all the numbers with 1's
func TestUpdateAllOnes(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	size := 100
	checkArr := make([]int, size)
	newArr := list.NewArrayList(size)

	for i := 0; i < size; i++ {
		err := newArr.Insert(newArr.Length(), i)
		if err != nil {
			t.Error(err)
		}
		checkArr[i] = i
	}

	for i := 0; i < newArr.Length(); i++ {
		val, err := newArr.Get(i)
		if err != nil {
			t.Error(err)
		}

		if val != checkArr[i] {
			t.Errorf("New ArrayList was inserted incorrectly %d != %d", val, checkArr[i])
		}
	}

	for i := 0; i < newArr.Length(); i++ {
		err := newArr.Update(i, 1)
		if err != nil {
			t.Error(err)
		}
	}

	for i := 0; i < newArr.Length(); i++ {
		val, err := newArr.Get(i)
		if err != nil {
			t.Error(err)
		}

		if val != 1 {
			t.Errorf("Index %d is not 1", i)
		}
	}
}
