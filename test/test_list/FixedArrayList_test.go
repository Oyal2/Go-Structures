package test_list

import (
	"github.com/oyal2/Go-Structures/list/FixedArrayList"
	"math/rand"
	"time"

	"testing"
)

//Test the ArrayList when its fully filled
func TestFillFixedArray(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(100)
	checkArr := make([]int, size)
	newArr := FixedArrayList.NewFixedArrayList(size)

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
func TestGetFixedArrayIndex(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(100)
	checkArr := make([]int, size)
	newArr := FixedArrayList.NewFixedArrayList(size)

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

//Test the ArrayList when inserted more than initialized
func TestOverFilledFixedArray(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(100)
	checkArr := make([]int, size)
	newArr := FixedArrayList.NewFixedArrayList(size)

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

	err := newArr.Insert(newArr.Length(), 0)
	if err == nil || err.Error() != "array capacity is full" {
		t.Errorf(err.Error())
	}
}

//Test to update all the numbers with 1's
func TestUpdateAllOnesFixedArray(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	size := 100
	checkArr := make([]int, size)
	newArr := FixedArrayList.NewFixedArrayList(size)

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
		err := newArr.Set(i, 1)
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
