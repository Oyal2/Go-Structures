package CircularQueue

import (
	"math/rand"
	"testing"
	"time"

	"github.com/oyal2/Go-Structures/list/ArrayList"
)

func TestCircularQueue(t *testing.T) {
	// create a new circular queue with an underlying array list and a maximum capacity of 5
	arrayList := ArrayList.New[int](5)
	cq := New[int](arrayList, 5)

	// the queue should initially be empty
	if cq.Length() != 0 {
		t.Errorf("Expected queue to be empty, but got length %d", cq.Length())
	}
	if !cq.IsEmpty() {
		t.Error("Expected queue to be empty, but it was not")
	}

	// enqueue some elements
	err := cq.Enqueue(1, 2, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if cq.Length() != 3 {
		t.Errorf("Expected queue to have length 3, but got %d", cq.Length())
	}

	// the front of the queue should be 1
	front, err := cq.Front()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if front != 1 {
		t.Errorf("Expected front of queue to be 1, but got %d", front)
	}

	// dequeue an element
	element, err := cq.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if element != 1 {
		t.Errorf("Expected dequeued element to be 1, but got %d", element)
	}
	if cq.Length() != 2 {
		t.Errorf("Expected queue to have length 2, but got %d", cq.Length())
	}

	// the front of the queue should now be 2
	front, err = cq.Front()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if front != 2 {
		t.Errorf("Expected front of queue to be 2, but got %d", front)
	}

	// try to enqueue more elements than the maximum capacity of the queue
	err = cq.Enqueue(4, 5, 6, 7)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err.Error() != "queue is full" {
		t.Errorf("Expected error message 'queue is full', but got '%s'", err.Error())
	}
	if cq.Length() != 2 {
		t.Errorf("Expected queue to have length 2, but got %d", cq.Length())
	}

	// clear the queue
	cq.Clear()
	if cq.Length() != 0 {
		t.Errorf("Expected queue to be empty, but got length %d", cq.Length())
	}
	if !cq.IsEmpty() {
		t.Error("Expected queue to be empty, but it was not")
	}

	// try to dequeue an element from an empty queue
	_, err = cq.Dequeue()
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err.Error() != "queue is empty" {
		t.Errorf("Expected error message 'queue is empty', but got '%s'", err.Error())
	}
}

func TestCircularQueueEnqueue(t *testing.T) {
	// create a new circular queue with an underlying array list and a maximum capacity of 5
	arrayList := ArrayList.New[int](5)
	cq := New[int](arrayList, 5)

	// enqueue some elements
	err := cq.Enqueue(1, 2, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if cq.Length() != 3 {
		t.Errorf("Expected queue to have length 3, but got %d", cq.Length())
	}

	// try to enqueue more elements than the maximum capacity of the queue
	err = cq.Enqueue(4, 5, 6)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err.Error() != "queue is full" {
		t.Errorf("Expected error message 'queue is full', but got '%s'", err.Error())
	}
	if cq.Length() != 3 {
		t.Errorf("Expected queue to have length 3, but got %d", cq.Length())
	}
}

func TestCircularQueueFront(t *testing.T) {
	// create a new circular queue with an underlying array list and a maximum capacity of 5
	arrayList := ArrayList.New[int](5)
	cq := New[int](arrayList, 5)

	// try to get the front element of an empty queue
	_, err := cq.Front()
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err.Error() != "queue is empty" {
		t.Errorf("Expected error message 'queue is empty', but got '%s'", err.Error())
	}

	// enqueue some elements
	err = cq.Enqueue(1, 2, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// the front of the queue should be 1
	front, err := cq.Front()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if front != 1 {
		t.Errorf("Expected front of queue to be 1, but got %d", front)
	}
}

func TestCircularQueueDequeue(t *testing.T) {
	// create a new circular queue with an underlying array list and a maximum capacity of 5
	arrayList := ArrayList.New[int](5)
	cq := New[int](arrayList, 5)

	// try to dequeue an element from an empty queue
	_, err := cq.Dequeue()
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err.Error() != "queue is empty" {
		t.Errorf("Expected error message 'queue is empty', but got '%s'", err.Error())
		// enqueue some elements
	}
	err = cq.Enqueue(1, 2, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// dequeue an element
	element, err := cq.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if element != 1 {
		t.Errorf("Expected dequeued element to be 1, but got %d", element)
	}
	if cq.Length() != 2 {
		t.Errorf("Expected queue to have length 2, but got %d", cq.Length())
	}

	// dequeue all remaining elements
	for !cq.IsEmpty() {
		_, err := cq.Dequeue()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}
	if cq.Length() != 0 {
		t.Errorf("Expected queue to be empty, but got length %d", cq.Length())
	}
}

func TestCircularQueueIsEmpty(t *testing.T) {
	// create a new circular queue with an underlying array list and a maximum capacity of 5
	arrayList := ArrayList.New[int](5)
	cq := New[int](arrayList, 5)

	// the queue should initially be empty
	if !cq.IsEmpty() {
		t.Error("Expected queue to be empty, but it was not")
	}

	// enqueue an element
	err := cq.Enqueue(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if cq.IsEmpty() {
		t.Error("Expected queue to be not empty, but it was")
	}

	// dequeue the element
	_, err = cq.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !cq.IsEmpty() {
		t.Error("Expected queue to be empty, but it was not")
	}
}

func TestCircularQueueLength(t *testing.T) {
	// create a new circular queue with an underlying array list and a maximum capacity of 5
	arrayList := ArrayList.New[int](5)
	cq := New[int](arrayList, 5)

	// the queue should initially be empty
	if cq.Length() != 0 {
		t.Errorf("Expected queue to have length 0, but got %d", cq.Length())
	}

	// enqueue an element
	err := cq.Enqueue(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if cq.Length() != 1 {
		t.Errorf("Expected queue to have length 1, but got %d", cq.Length())
	}
	// clear the queue
	cq.Clear()
	if cq.Length() != 0 {
		t.Errorf("Expected queue to have length 0, but got %d", cq.Length())
	}
}

func TestCircularQueueEnqueueLargeRandomInserts(t *testing.T) {
	// create a new circular queue with an underlying array list and a maximum capacity of 1000
	arrayList := ArrayList.New[int](1000)
	cq := New[int](arrayList, 1000)

	// create a random number generator
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// enqueue 1000 random elements
	for i := 0; i < 1000; i++ {
		err := cq.Enqueue(rng.Int())
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}
	if cq.Length() != 1000 {
		t.Errorf("Expected queue to have length 1000, but got %d", cq.Length())
	}

	// try to enqueue one more element
	err := cq.Enqueue(rng.Int())
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err.Error() != "queue is full" {
		t.Errorf("Expected error message 'queue is full', but got '%s'", err.Error())
	}
}

func TestCircularQueueDequeueLargeRandomInserts(t *testing.T) {
	// create a new circular queue with an underlying array list and a maximum capacity of 1000
	arrayList := ArrayList.New[int](1000)
	cq := New[int](arrayList, 1000)

	// create a random number generator
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// enqueue 1000 random elements
	for i := 0; i < 1000; i++ {
		err := cq.Enqueue(rng.Int())
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}

	// dequeue 1000 elements
	for i := 0; i < 1000; i++ {
		_, err := cq.Dequeue()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}
	if cq.Length() != 0 {
		t.Errorf("Expected queue to have length 0, but got %d", cq.Length())
	}

	// try to dequeue one more element
	_, err := cq.Dequeue()
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err.Error() != "queue is empty" {
		t.Errorf("Expected error message 'queue is empty', but got '%s'", err.Error())
	}
}
