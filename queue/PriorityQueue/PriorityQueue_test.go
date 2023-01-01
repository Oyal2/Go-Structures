package PriorityQueue

import (
	"testing"

	"github.com/oyal2/Go-Structures/heap/BinaryHeap"
	"github.com/oyal2/Go-Structures/heap/FibonacciHeap"
)

func TestNew(t *testing.T) {
	h := BinaryHeap.New(func(a, b int) bool {
		return a < b
	})
	q := New[int](h)
	if q._storage != h {
		t.Errorf("Error in New(): Incorrect storage")
	}
}

func TestEnqueue(t *testing.T) {
	h := BinaryHeap.New(func(a, b int) bool {
		return a < b
	})
	q := New[int](h)
	err := q.Enqueue(1, 2, 3)
	if err != nil {
		t.Errorf("Error in Enqueue(): %s", err)
	}
	if q.Length() != 3 {
		t.Errorf("Error in Enqueue(): Incorrect number of elements in queue")
	}
}

func TestFront(t *testing.T) {
	h := BinaryHeap.New(func(a, b int) bool {
		return a < b
	})
	q := New[int](h)
	q.Enqueue(3, 2, 1)
	elem, err := q.Front()
	if err != nil {
		t.Errorf("Error in Front(): %s", err)
	}
	if elem != 1 {
		t.Errorf("Error in Front(): Incorrect element returned")
	}
}

func TestDequeue(t *testing.T) {
	h := BinaryHeap.New(func(a, b int) bool {
		return a < b
	})
	q := New[int](h)
	q.Enqueue(3, 2, 1)
	elem, err := q.Dequeue()
	if err != nil {
		t.Errorf("Error in Dequeue(): %s", err)
	}
	if elem != 1 {
		t.Errorf("Error in Dequeue(): Incorrect element returned")
	}
	if q.Length() != 2 {
		t.Errorf("Error in Dequeue(): Incorrect number of elements in queue")
	}
}

func TestContains(t *testing.T) {
	h := BinaryHeap.New(func(a, b int) bool {
		return a < b
	})
	q := New[int](h)
	q.Enqueue(1, 2, 3)
	if !q.Contains(2) {
		t.Errorf("Error in Contains(): Element not found in queue")
	}
	if q.Contains(4) {
		t.Errorf("Error in Contains(): Element found in queue when it should not have been")
	}
}

func TestLength(t *testing.T) {
	h := BinaryHeap.New(func(a, b int) bool {
		return a < b
	})
	q := New[int](h)
	q.Enqueue(1, 2, 3)
	if q.Length() != 3 {
		t.Errorf("Error in Length(): Incorrect length of queue")
	}
}

func TestClear(t *testing.T) {
	h := BinaryHeap.New(func(a, b int) bool {
		return a < b
	})
	q := New[int](h)
	q.Enqueue(1, 2, 3)
	q.Clear()
	if q.Length() != 0 {
		t.Errorf("Error in Clear(): Queue not empty after clear")
	}
}

func TestUpdate(t *testing.T) {
	h := BinaryHeap.New(func(a, b int) bool {
		return a < b
	})
	q := New[int](h)
	q.Enqueue(1, 2, 3)
	err := q.Update(2, 4)
	if err != nil {
		t.Errorf("Error in Update(): %s", err)
	}
	if q.Length() != 3 {
		t.Errorf("Error in Update(): Incorrect number of elements in queue")
	}
	elem, err := q.Front()
	if err != nil {
		t.Errorf("Error in Update(): %s", err)
	}
	if elem != 1 {
		t.Errorf("Error in Update(): Incorrect element at front of queue")
	}
}

func TestIsEmpty(t *testing.T) {
	h := BinaryHeap.New(func(a, b int) bool {
		return a < b
	})
	q := New[int](h)
	if !q.IsEmpty() {
		t.Errorf("Error in IsEmpty(): Queue is not empty")
	}
	q.Enqueue(1)
	if q.IsEmpty() {
		t.Errorf("Error in IsEmpty(): Queue is empty")
	}
}

func TestNewFibonacci(t *testing.T) {
	h := FibonacciHeap.New(func(a, b int) bool {
		return a < b
	})
	q := New[int](h)
	if q._storage != h {
		t.Errorf("Error in New(): Incorrect storage")
	}
}

func TestEnqueueFibonacci(t *testing.T) {
	h := FibonacciHeap.New(func(a, b int) bool {
		return a < b
	})
	q := New[int](h)
	err := q.Enqueue(1, 2, 3)
	if err != nil {
		t.Errorf("Error in Enqueue(): %s", err)
	}
	if q.Length() != 3 {
		t.Errorf("Error in Enqueue(): Incorrect number of elements in queue")
	}
}

func TestFrontFibonacci(t *testing.T) {
	h := FibonacciHeap.New(func(a, b int) bool {
		return a < b
	})
	q := New[int](h)
	q.Enqueue(3, 2, 1)
	elem, err := q.Front()
	if err != nil {
		t.Errorf("Error in Front(): %s", err)
	}
	if elem != 1 {
		t.Errorf("Error in Front(): Incorrect element returned")
	}
}

func TestDequeueFibonacci(t *testing.T) {
	h := FibonacciHeap.New(func(a, b int) bool {
		return a < b
	})
	q := New[int](h)
	q.Enqueue(3, 2, 1)
	elem, err := q.Dequeue()
	if err != nil {
		t.Errorf("Error in Dequeue(): %s", err)
	}
	if elem != 1 {
		t.Errorf("Error in Dequeue(): Incorrect element returned")
	}
	if q.Length() != 2 {
		t.Errorf("Error in Dequeue(): Incorrect number of elements in queue")
	}
}

func TestContainsFibonacci(t *testing.T) {
	h := FibonacciHeap.New(func(a, b int) bool {
		return a < b
	})
	q := New[int](h)
	q.Enqueue(1, 2, 3)
	if !q.Contains(2) {
		t.Errorf("Error in Contains(): Element not found in queue")
	}
	if q.Contains(4) {
		t.Errorf("Error in Contains(): Element found in queue when it should not have been")
	}
}

func TestIsEmptyFibonacci(t *testing.T) {
	h := FibonacciHeap.New(func(a, b int) bool {
		return a < b
	})
	q := New[int](h)
	if !q.IsEmpty() {
		t.Errorf("Error in IsEmpty(): Queue is not empty")
	}
	q.Enqueue(1)
	if q.IsEmpty() {
		t.Errorf("Error in IsEmpty(): Queue is empty")
	}
}

func TestLengthFibonacci(t *testing.T) {
	h := FibonacciHeap.New(func(a, b int) bool {
		return a < b
	})
	q := New[int](h)
	q.Enqueue(1, 2, 3)
	if q.Length() != 3 {
		t.Errorf("Error in Length(): Incorrect length of queue")
	}
}

func TestClearFibonacci(t *testing.T) {
	h := FibonacciHeap.New(func(a, b int) bool {
		return a < b
	})
	q := New[int](h)
	q.Enqueue(1, 2, 3)
	q.Clear()
	if q.Length() != 0 {
		t.Errorf("Error in Clear(): Queue not empty after clear")
	}
}

func TestUpdateFibonacci(t *testing.T) {
	h := FibonacciHeap.New(func(a, b int) bool {
		return a < b
	})
	q := New[int](h)
	q.Enqueue(1, 2, 3)
	err := q.Update(3, 0)
	if err != nil {
		t.Errorf("Error in Update(): %s", err)
	}
	if q.Length() != 3 {
		t.Errorf("Error in Update(): Incorrect number of elements in queue")
	}
	elem, err := q.Front()
	if err != nil {
		t.Errorf("Error in Update(): %s", err)
	}
	if elem != 0 {
		t.Errorf("Error in Update(): Incorrect element at front of queue")
	}
}
