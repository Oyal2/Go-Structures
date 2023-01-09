package BinaryHeap

import (
	"math/rand"
	"sort"
	"testing"
)

func TestInsert(t *testing.T) {
	nums := []int{1, 3, 6, 5, 9, 8}
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	binaryHeap := New(func(a, b int) bool {
		return a < b
	})

	for _, val := range nums {
		binaryHeap.Insert(val)
	}

	if binaryHeap.storage[1] != 1 {
		t.Errorf("Did not get the min value")
	}
}

func TestCrazyInsert(t *testing.T) {
	var nums []int

	for i := 0; i < rand.Intn(10000); i++ {
		nums = append(nums, rand.Intn(10000))
	}
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	binaryHeap := New(func(a, b int) bool {
		return a < b
	})

	for _, val := range nums {
		binaryHeap.Insert(val)
	}

	if binaryHeap.storage[1] != sortedNums[0] {
		t.Errorf("Did not get the min value")
	}
}

func TestMin(t *testing.T) {
	nums := []int{1, 3, 6, 5, 9, 8}
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	binaryHeap := New(func(a, b int) bool {
		return a < b
	})

	for _, val := range nums {
		binaryHeap.Insert(val)
	}

	var val int
	var err error
	for count := 0; count < len(sortedNums); count++ {
		val, err = binaryHeap.Extract()
		if err != nil {
			t.Error(err)
		} else if val != sortedNums[count] {
			t.Errorf("Minimum was taken incorrectly. Answer: %d vs Output: %d", sortedNums[count], val)
		}
	}
	_, err = binaryHeap.Peek()
	if err == nil {
		t.Error("error should be thrown")
	}
}

func TestCrazyMin(t *testing.T) {
	var nums []int

	for i := 0; i < rand.Intn(1000); i++ {
		nums = append(nums, rand.Intn(1000))
	}
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	binaryHeap := New(func(a, b int) bool {
		return a < b
	})

	for _, val := range nums {
		binaryHeap.Insert(val)
	}

	var val int
	var err error
	for count := 0; count < len(sortedNums); count++ {
		val, err = binaryHeap.Extract()
		if err != nil {
			t.Error(err)
		} else if val != sortedNums[count] {
			t.Errorf("Minimum was taken incorrectly. Answer: %d vs Output: %d", sortedNums[count], val)
		}
	}
	_, err = binaryHeap.Peek()
	if err == nil {
		t.Error("error should be thrown")
	}
}

func TestPeek(t *testing.T) {
	nums := []int{1, 3, 6, 5, 9, 8}
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	binaryHeap := New(func(a, b int) bool {
		return a < b
	})

	for _, val := range nums {
		binaryHeap.Insert(val)
	}

	var val int
	var err error
	for count := 0; count < len(sortedNums); count++ {
		val, err = binaryHeap.Peek()
		if err != nil {
			t.Error(err)
		} else if val != sortedNums[count] {
			t.Errorf("Minimum was taken incorrectly. Answer: %d vs Output: %d", sortedNums[count], val)
		}
		binaryHeap.Extract()
	}
	_, err = binaryHeap.Peek()
	if err == nil {
		t.Error("error should be thrown")
	}
}

func TestLength(t *testing.T) {
	nums := []int{1, 3, 6, 5, 9, 8}
	binaryHeap := New(func(a, b int) bool {
		return a < b
	})

	for _, val := range nums {
		binaryHeap.Insert(val)
	}

	if binaryHeap.Length() != len(nums) {
		t.Error("incorrect size")
	}
}

func TestContains(t *testing.T) {
	nums := []int{1, 3, 6, 5, 9, 8}
	binaryHeap := New(func(a, b int) bool {
		return a < b
	})

	for _, val := range nums {
		binaryHeap.Insert(val)
	}

	if !binaryHeap.Contains(3) {
		t.Error("3 does exist in heap")
	}

	if binaryHeap.Contains(20) {
		t.Error("20 doesnt exist in heap")
	}
}

func TestDecreaseKey(t *testing.T) {
	nums := []int{1, 3, 6, 5, 9, 8}
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	binaryHeap := New(func(a, b int) bool {
		return a < b
	})

	for _, val := range nums {
		binaryHeap.Insert(val)
	}

	val, err := binaryHeap.Peek()
	if err != nil {
		t.Error(err)
	}

	if val != 1 {
		t.Error("Value should be 1, got ", val)
	}

	err = binaryHeap.ChangeKey(1, 10)
	if err != nil {
		t.Error(err)
	}

	val, err = binaryHeap.Peek()

	if err != nil {
		t.Error(err)
	}

	if val != 3 {
		t.Error("Value should be 3, got ", val)
	}

	newNums := []int{3, 5, 6, 8, 9, 10}
	for count := 0; count < len(newNums); count++ {
		val, err = binaryHeap.Extract()
		if err != nil {
			t.Error(err)
		} else if val != newNums[count] {
			t.Errorf("Minimum was taken incorrectly. Answer: %d vs Output: %d", newNums[count], val)
		}
	}
}

func TestDecreaseKeyCrazy(t *testing.T) {
	var nums []int
	for i := 0; i < rand.Intn(1000); i++ {
		nums = append(nums, rand.Intn(1000))
	}
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	binaryHeap := New(func(a, b int) bool {
		return a < b
	})

	for _, val := range nums {
		binaryHeap.Insert(val)
	}

	val, err := binaryHeap.Peek()
	if err != nil {
		t.Error(err)
	}

	if val != sortedNums[0] {
		t.Error("Value should be 1, got ", val)
	}

	newNum := rand.Intn(1000)
	indexNum := rand.Intn(len(nums))
	err = binaryHeap.ChangeKey(binaryHeap.storage[indexNum], newNum)
	if err != nil {
		t.Error(err)
	}

	sortedNums[indexNum] = newNum
	sort.Ints(sortedNums)
	for count := 0; count < len(sortedNums); count++ {
		val, err = binaryHeap.Extract()
		if err != nil {
			t.Error(err)
		} else if val != sortedNums[count] {
			t.Errorf("Minimum was taken incorrectly. Answer: %d vs Output: %d", sortedNums[count], val)
		}
	}
}
