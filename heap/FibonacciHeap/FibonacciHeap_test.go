package FibonacciHeap

import (
	"math"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestExtract(t *testing.T) {
	// create a new Fibonacci heap
	h := New(func(a, b int) bool {
		return a < b
	})

	// insert some elements into the heap
	h.Insert(5)
	h.Insert(3)
	h.Insert(7)
	h.Insert(2)

	// extract the minimum element
	min, err := h.Extract()
	if err != nil {
		t.Errorf("Error extracting minimum element: %v", err)
	}
	if min != 2 {
		t.Errorf("Expected minimum element to be 2, got %v", min)
	}

	h.Display()
	// extract the remaining elements
	/**
		7 <--> 3
			   ^
			   |
			   v
			   5
	**/
	min, err = h.Extract()
	if err != nil {
		t.Errorf("Error extracting minimum element: %v", err)
	}
	if min != 3 {
		t.Errorf("Expected minimum element to be 3, got %v", min)
	}
	h.Display()
	// [5,7]
	min, err = h.Extract()
	if err != nil {
		t.Errorf("Error extracting minimum element: %v", err)
	}
	if min != 5 {
		t.Errorf("Expected minimum element to be 5, got %v", min)
	}
	// [7]
	min, err = h.Extract()
	if err != nil {
		t.Errorf("Error extracting minimum element: %v", err)
	}
	if min != 7 {
		t.Errorf("Expected minimum element to be 7, got %v", min)
	}

	// try to extract from an empty heap
	min, err = h.Extract()
	if err == nil {
		t.Errorf("Expected error extracting from empty heap, got %v", min)
	}
}

func TestInsert(t *testing.T) {
	// create a new Fibonacci heap
	h := New(func(a, b int) bool {
		return a < b
	})

	// insert some elements into the heap
	err := h.Insert(5, 3, 7, 2)
	if err != nil {
		t.Errorf("Error inserting elements into heap: %v", err)
	}

	// check the size of the heap
	if h.Length() != 4 {
		t.Errorf("Expected heap size to be 4, got %v", h.Length())
	}

	// check the minimum element in the heap
	min, err := h.Peek()
	if err != nil {
		t.Errorf("Error finding minimum element in heap: %v", err)
	}
	if min != 2 {
		t.Errorf("Expected minimum element to be 2, got %v", min)
	}
}

func TestInsert2(t *testing.T) {
	// create a new Fibonacci heap
	h := New(func(a, b int) bool {
		return a < b
	})
	rand.Seed(time.Now().UnixNano())
	// insert a large number of randomly generated elements into the heap
	n := 7
	nums := []int{13,
		62,
		31,
		58,
		24,
		94,
		12}
	for i := 0; i < n; i++ {
		// num := rand.Intn(100)
		err := h.Insert(nums[i])
		if err != nil {
			t.Errorf("Error inserting elements into heap: %v", err)
		}
		// println(nums[i])
	}
	// extract the minimum element from the heap repeatedly
	prevMin := math.MinInt32
	for i := 0; i < n; i++ {
		min, err := h.Extract()
		if err != nil {
			t.Errorf("Error extracting minimum element: %v", err)
		}
		if min < prevMin {
			t.Errorf("Expected minimum element to be smaller than previous minimum %d > %d", min, prevMin)
		}
		prevMin = min
	}
}

func TestRandomizedInsert(t *testing.T) {
	// create a new Fibonacci heap
	h := New(func(a, b int) bool {
		return a < b
	})
	rand.Seed(time.Now().UnixNano())
	// insert a large number of randomly generated elements into the heap
	n := 1000
	for i := 0; i < n; i++ {
		num := rand.Int()
		err := h.Insert(num)
		if err != nil {
			t.Errorf("Error inserting elements into heap: %v", err)
		}
	}
	// extract the minimum element from the heap repeatedly
	prevMin := math.MinInt32
	for i := 0; i < n; i++ {
		min, err := h.Extract()
		if err != nil {
			t.Errorf("Error extracting minimum element: %v", err)
		}
		if min < prevMin {
			t.Errorf("Expected minimum element to be smaller than previous minimum %d > %d", min, prevMin)
		}
		prevMin = min
	}
}

func TestExtractRandomized(t *testing.T) {
	// create a new Fibonacci heap
	h := New(func(a, b int) bool {
		return a < b
	})

	// generate 100 random elements and insert them into the heap
	rand.Seed(time.Now().UnixNano())
	expected := make([]int, 100)
	for i := 0; i < 100; i++ {
		element := rand.Intn(1000)
		h.Insert(element)
		expected[i] = element
	}

	// sort the expected values in ascending order
	sort.Ints(expected)

	// extract the elements from the heap and check that they are in ascending order
	for i := 0; i < 100; i++ {
		min, err := h.Extract()
		if err != nil {
			t.Errorf("Error extracting minimum element: %v", err)
		}
		if min != expected[i] {
			t.Errorf("Expected minimum element to be %v, got %v", expected[i], min)
		}
	}
}

func TestClear(t *testing.T) {
	// create a new Fibonacci heap
	h := New(func(a, b int) bool {
		return a < b
	})

	// insert some elements into the heap
	h.Insert(5)
	h.Insert(3)
	h.Insert(7)
	h.Insert(2)

	// clear the heap
	h.Clear()

	// check that the heap is empty
	if h.Length() != 0 {
		t.Errorf("Expected heap to be empty, got size %v", h.Length())
	}

	// check that extracting from an empty heap returns an error
	_, err := h.Extract()
	if err == nil {
		t.Errorf("Expected error extracting from empty heap, got no error")
	}
}

func TestPeek(t *testing.T) {
	// create a new Fibonacci heap
	h := New(func(a, b int) bool {
		return a < b
	})

	// insert some elements into the heap
	h.Insert(5)
	h.Insert(3)
	h.Insert(7)
	h.Insert(2)

	// check that the minimum element can be retrieved
	min, err := h.Peek()
	if err != nil {
		t.Errorf("Error retrieving minimum element: %v", err)
	}
	if min != 2 {
		t.Errorf("Expected minimum element to be 2, got %v", min)
	}

	// check that the heap size is unchanged
	if h.Length() != 4 {
		t.Errorf("Expected heap size to be 4, got %v", h.Length())
	}
}

func TestIsEmptyAndLength(t *testing.T) {
	// create a new Fibonacci heap
	h := New(func(a, b int) bool {
		return a < b
	})

	// check that the heap is empty
	if !h.IsEmpty() {
		t.Errorf("Expected heap to be empty, got size %v", h.Length())
	}

	// insert some elements into the heap
	h.Insert(5)
	h.Insert(3)
	h.Insert(7)
	h.Insert(2)

	// check that the heap is not empty
	if h.IsEmpty() {
		t.Errorf("Expected heap to be non-empty, got size %v", h.Length())
	}

	// check that the heap size is correct
	if h.Length() != 4 {
		t.Errorf("Expected heap size to be 4, got %v", h.Length())
	}
}

func TestFibonacciHeap(t *testing.T) {
	// create a new Fibonacci heap
	h := New(func(a, b int) bool {
		return a < b
	})

	// insert some elements into the heap
	h.Insert(5)
	h.Insert(3)
	h.Insert(7)
	h.Insert(2)

	// check that the heap is not empty
	if h.IsEmpty() {
		t.Errorf("Expected heap to be non-empty, got size %v", h.Length())
	}

	// check that the heap size is correct
	if h.Length() != 4 {
		t.Errorf("Expected heap size to be 4, got %v", h.Length())
	}

	// check that the minimum element can be retrieved
	min, err := h.Peek()
	if err != nil {
		t.Errorf("Error retrieving minimum element: %v", err)
	}
	if min != 2 {
		t.Errorf("Expected minimum element to be 2, got %v", min)
	}

	// check that the heap size is unchanged
	if h.Length() != 4 {
		t.Errorf("Expected heap size to be 4, got %v", h.Length())
	}

	// extract the minimum element
	min, err = h.Extract()
	if err != nil {
		t.Errorf("Error extracting minimum element: %v", err)
	}
	if min != 2 {
		t.Errorf("Expected minimum element to be 2, got %v", min)
	}

	// check that the heap size is correct
	if h.Length() != 3 {
		t.Errorf("Expected heap size to be 3, got %v", h.Length())
	}

	// extract the remaining elements
	min, err = h.Extract()
	if err != nil {
		t.Errorf("Error extracting minimum element: %v", err)
	}
	if min != 3 {
		t.Errorf("Expected minimum element to be 3, got %v", min)
	}

	// check that the heap size is correct
	if h.Length() != 2 {
		t.Errorf("Expected heap size to be 2, got %v", h.Length())
	}

	min, err = h.Extract()
	if err != nil {
		t.Errorf("Error extracting minimum element: %v", err)
	}
	if min != 5 {
		t.Errorf("Expected minimum element to be 5, got %v", min)
	}

	// check that the heap size is correct
	if h.Length() != 1 {
		t.Errorf("Expected heap size to be 1, got %v", h.Length())
	}

	min, err = h.Extract()
	if err != nil {
		t.Errorf("Error extracting minimum element: %v", err)
	}
	if min != 7 {
		t.Errorf("Expected minimum element to be 7, got %v", min)
	}

	// check that the heap is now empty
	if !h.IsEmpty() {
		t.Errorf("Expected heap to be empty, got size %v", h.Length())
	}

	// try to extract from an empty heap
	min, err = h.Extract()
	if err == nil {
		t.Errorf("Expected error extracting from empty heap, got %v", min)
	}

	// clear the heap
	h.Clear()

	// check that the heap is now empty
	if !h.IsEmpty() {
		t.Errorf("Expected heap to be empty, got size %v", h.Length())
	}
}

func TestChangeKey(t *testing.T) {
	// create a new Fibonacci heap
	h := New(func(a, b int) bool {
		return a < b
	})

	// insert some elements into the heap
	h.Insert(5, 3, 7, 2)

	// change the key of an element
	err := h.ChangeKey(7, 1)
	if err != nil {
		t.Errorf("Error changing key: %v", err)
	}

	// check that the minimum element is now 1
	min, err := h.Extract()
	if err != nil {
		t.Errorf("Error extracting minimum element: %v", err)
	}
	if min != 1 {
		t.Errorf("Expected minimum element to be 1, got %v", min)
	}
}

func TestContains(t *testing.T) {
	// create a new Fibonacci heap
	h := New(func(a, b int) bool {
		return a < b
	})

	// insert some elements into the heap
	h.Insert(5)
	h.Insert(3)
	h.Insert(7)
	h.Insert(2)

	// check if the heap contains certain elements
	if !h.Contains(5) {
		t.Errorf("Expected heap to contain 5, got false")
	}
	if !h.Contains(3) {
		t.Errorf("Expected heap to contain 3, got false")
	}
	if !h.Contains(7) {
		t.Errorf("Expected heap to contain 7, got false")
	}
	if !h.Contains(2) {
		t.Errorf("Expected heap to contain 2, got false")
	}
	if h.Contains(1) {
		t.Errorf("Expected heap not to contain 1, got true")
	}
	if h.Contains(4) {
		t.Errorf("Expected heap not to contain 4, got true")
	}
	if h.Contains(6) {
		t.Errorf("Expected heap not to contain 6, got true")
	}
	if h.Contains(8) {
		t.Errorf("Expected heap not to contain 8, got true")
	}
}
