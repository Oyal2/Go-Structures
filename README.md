<p align="center"><img  src="https://user-images.githubusercontent.com/13637813/193483325-7b8b64c4-577d-43e1-a47c-ee0b75eb5bd0.png" width=400px>
</p>
<h2 align="center" style="border-bottom: none">Go Structures</h2>
<div align="center">

  [![Status](https://img.shields.io/badge/status-active-success.svg)]()
   [![Go Report Card](https://goreportcard.com/badge/github.com/oyal2/Go-Structures)](https://goreportcard.com/report/github.com/oyal2/Go-Structures)
  ![GitHub Pull Requests](https://img.shields.io/github/go-mod/go-version/oyal2/Go-Structures)
  [![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)
</div>

---

## 📝 Table of Contents
- [📝 Table of Contents](#-table-of-contents)
- [🧐 About ](#-about-)
- [⚙️ Data Structures ](#️-data-structures-)
	- [List](#list)
	- [Tree](#tree)
	- [Hashed](#hashed)
	- [Probabilistic](#probabilistic)
	- [Heap](#heap)
	- [Queue](#queue)
	- [Stack](#stack)
- [🎈 Usage ](#-usage-)
	- [Sort Interface](#sort-interface)
	- [List](#list-1)
		- [ArrayList](#arraylist-1)
		- [Fixed ArrayList](#fixed-arraylist)
	- [Node Type](#node-type)
		- [Singly Linked List](#singly-linked-list)
		- [Doubly Linked List](#doubly-linked-list)
	- [Queue](#queue-1)
		- [Generic Queue](#generic-queue)
		- [Circular Queue](#circular-queue)
		- [Priority Queue](#priority-queue)
	- [Stack](#stack-1)
		- [Generic Stack](#generic-stack)
	- [Probabilistic](#probabilistic-1)
		- [Bloom Filter](#bloom-filter)
	- [Heap](#heap-1)
		- [Binary Heap](#binary-heap)
		- [Fibonacci Heap](#fibonacci-heap)

## 🧐 About <a name = "about"></a>
I made this library because Golang has a very limited amount of data structures. I implemented various data structures such as ArrayList, Queue, DoublyLinked List, SinglyLinked List, and many more. There are many complex data structures in this library and some basic ones. All Golang data structures are optimized and are generic classes. 

## ⚙️ Data Structures <a name = "data_structures"></a>
### List

- [x] ArrayList
- [x] Singly Linked List
- [x] Doubly Linked List
- [x] Fixed Array List


### Tree
- [ ] Van Emde Boas Trees
- [x] AVL Tree
- [ ] B-Tree
- [ ] Trei
- [ ] Merkel Trees
- [ ] Log-Structured Merge-Tree

### Hashed
- [ ] Hash Table
- [ ] Ctrie

### Probabilistic
- [x] Bloom Filter
- [ ] Skip List
- [ ] HyperLogLog
- [ ] Count–Min Sketch
- [ ] Quotient Filter

### Heap
- [X] Binary Heap
- [X] Fibonacci Heap

### Queue
- [x] Queue
- [x] Circular Queue
- [x] Priority Queue

### Stack
- [x] Stack

## 🎈 Usage <a name = "usage"></a>

### Sort Interface 

```go 
type Sort[T Ordered] interface {
	Less(a, b T) bool
	Swap(i, j int)
}
```

### List

```go
type List[T comparable] interface {
	Get(index int) (T, error)
	Insert(index int, element T) error
	Remove(index int) (T, error)
	Length() int
    Update(index int, element T) error
}
```

#### ArrayList

A [list](https://en.wikipedia.org/wiki/Dynamic_array) that dynamically grows by n*2.
```go
package main

import (
    "github.com/oyal2/Go-Structures/list/ArrayList"
)

func main() {
	list := ArrayList.New[int](3)
	list.Add(1)                             // [1]
	list.Add(2,3)                           // [1,2,3]
	_, _ = list.Get(0)                      // 1,nil
	_, _ = list.Get(100)                    // nil,the index is out of bounds
	_ = list.Contains(1)                    // true
	_ = list.Contains(4)                    // false
	_,_ = list.Remove(2)                    // 3,nil
	_,_ = list.Remove(1)                    // 2,nil
	_,_ = list.Remove(0)                    // 1,nil
	_,_ = list.Remove(0)                    // nil,index 0 is out of bounds
	_ = list.Empty()                        // true
	_ = list.Length()                       // 0
	list.Add(1)                             // [1]
	list.Clear()                            // []
	list.Insert(0, 2)                       // [2]
	list.Insert(0, 1)                       // [1,2]
}
```

#### Fixed ArrayList

An immutable ArrayList
```go
package main

import (
    "github.com/oyal2/Go-Structures/list/FixedArrayList"
)

func main() {
	list := FixedArrayList.New[int](2)
	list.Insert(0,1)                        // [1]
	list.Insert(1,2)                        // [1,2]
    list.Insert(2)                          // array capacity is full
	_, _ = list.Get(0)                      // 1,nil
	_, _ = list.Get(100)                    // nil,the index is out of bounds
	_ = list.Contains(1)                    // true
	_ = list.Contains(4)                    // false
	_,_ = list.Remove(1)                    // 2,nil
	_,_ = list.Remove(0)                    // 1,nil
	_,_ = list.Remove(0)                    // nil,index 0 is out of bounds
	_ = list.Empty()                        // true
	_ = list.Length()                       // 0
	list.Add(1)                             // [1]
	list.Clear()                            // []
}
```

### Node Type

```go
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}
```


#### Singly Linked List

A [List](https://en.wikipedia.org/wiki/Linked_list#Singly_linked_list) where each [node](https://github.com/Oyal2/Go-Structures#node-type) contains an element and points to the next node
```go
package main

import (
    "github.com/oyal2/Go-Structures/list/SinglyLinkedList"
)

func main() {
	list := SinglyLinkedList.New[int]()
	list.Add(1)                             // [1]
	list.Add(2,3)                           // [1,2,3]
	list.Prepend(4)                         // [4,1,2,3]
    list.Append(5)                          // [4,1,2,3,5]
	_, _ = list.Get(0)                      // 4,nil
	_, _ = list.Get(100)                    // nil,the index is out of bounds
	_ = list.Contains(1)                    // true
	_ = list.Contains(6)                    // false
	list.Remove(4)                          // 5,nil
	list.Remove(3)                          // 3,nil
	list.Remove(2)                          // 2,nil
	list.Remove(1)                          // 1,nil
	list.Remove(0)                          // 4,nil
	list.Remove(0)                          // nil,the index is out of bounds
	_ = list.Empty()                        // true
	_ = list.Length()                       // 0
	list.Add(1)                             // [1]
	list.Clear()                            // []
	list.Insert(0, 2)                       // [2]
	list.Insert(0, 1)                       // [1,2]
    list.Update(0, 3)                       // [3,2]
    list.IndexOf(2)                         // 1
    list.IndexOf(7)                         // -1
}
```


#### Doubly Linked List

A [List](https://en.wikipedia.org/wiki/Doubly_linked_list) where each [node](https://github.com/Oyal2/Go-Structures#node-type) contains an element and points to the next node and the preceeding node.
```go
package main

import (
    "github.com/oyal2/Go-Structures/list/DoublyLinkedList"
)

func main() {
	list := DoublyLinkedList.New[int]()
	list.Add(1)                             // [1]
	list.Add(2,3)                           // [1,2,3]
	list.Prepend(4)                         // [4,1,2,3]
    list.Append(5)                          // [4,1,2,3,5]
	_, _ = list.Get(0)                      // 4,nil
	_, _ = list.Get(100)                    // nil,the index is out of bounds
	_ = list.Contains(1)                    // true
	_ = list.Contains(6)                    // false
	list.Remove(4)                          // 5,nil
	list.Remove(3)                          // 3,nil
	list.Remove(2)                          // 2,nil
	list.Remove(1)                          // 1,nil
	list.Remove(0)                          // 4,nil
	list.Remove(0)                          // nil,the index is out of bounds
	_ = list.Empty()                        // true
	_ = list.Length()                       // 0
	list.Add(1)                             // [1]
	list.Clear()                            // []
	list.Insert(0, 2)                       // [2]
	list.Insert(0, 1)                       // [1,2]
    list.Update(0, 3)                       // [3,2]
    list.IndexOf(2)                         // 1
    list.IndexOf(7)                         // -1
}
```

### Queue


#### Generic Queue

A [Queue](https://en.wikipedia.org/wiki/Queue_(abstract_data_type)) is data structure that maintains a sequence of first-in-first-out (FIFO). Typical operations are `enqueue`, `dequeue`, and `front`

```go
type Queue[T comparable] struct {
	_list list.List[T]
}
```

```go
package main

import (
	"github.com/oyal2/Go-Structures/list/ArrayList"
	"github.com/oyal2/Go-Structures/queue/Queue"
)

func main() {
	queue := Queue.New[int](ArrayList.New[int](10))
    queue.Enqueue(1)                        // 1
    queue.Enqueue(2)                        // 1, 2
    _, _ = queue.Front()                    // 1,nil
    _, _ = queue.Dequeue()                  // 1, nil
    _, _ = queue.Dequeue()                  // 2, nil
    _, _ = queue.Dequeue()                  // nil, index is out of bounds
    queue.Enqueue(1)                        // 1
    queue.Clear()                           // empty
    queue.IsEmpty()                         // true
    _ = queue.Length()                      // 0
}
```

#### Circular Queue
A [Circular Queue](https://en.wikipedia.org/wiki/Circular_buffer) is a queue data structure that maintains a sequence of first-in-first-out (FIFO). It is implemented using a circular buffer, which allows it to be efficient in both time and space. Typical operations are `enqueue`, `dequeue`, and `front`.

```go
type CircularQueue[T comparable] struct {
	_storage list.List[T]
	_max     int
	_front   int
	_rear    int
	_size    int
}
```

```go

package main

import (
	"fmt"

	"github.com/oyal2/Go-Structures/list/ArrayList"
	"github.com/oyal2/Go-Structures/queue/CircularQueue"
)

func main() {
	queue := CircularQueue.New[int](ArrayList.New[int](5), 5)

    _ = cqueue.Enqueue(1)                    // 1
    _ = cqueue.Enqueue(2)                    // 1, 2
    _, _ = cqueue.Front()                    // 1,nil
    _, _ = cqueue.Dequeue()                  // 1, nil
    _, _ = cqueue.Dequeue()                  // 2, nil
    _, _ = cqueue.Dequeue()                  // nil, queue is empty
    cqueue.Enqueue(1)                        // 1
    cqueue.Clear()                           // empty
    cqueue.IsEmpty()                         // true
    _ = cqueue.Length()                      // 0
	_ = cqueue.Enqueue(5,1,2,4,6,9)          // queue is full
```


#### Priority Queue

A [Priority Queue](https://en.wikipedia.org/wiki/Priority_queue) is data structure very similar to a queue, but instead of respecting FIFO, an element with the highest prioirty is dequeued. Typical operations are `enqueue`, `dequeue`, and `front`. The storage type for this queue is only supported with heaps.
```go
type PriorityQueue[T utils.Ordered] struct {
	_storage heap.Heap[T]
}
```


```go
package main

import (
	"github.com/oyal2/Go-Structures/heap/BinaryHeap"
	"github.com/oyal2/Go-Structures/queue/PriorityQueue"
)

func main() {
	comparator := func(a, b int) bool {
		return a < b
	}
	binaryHeap := BinaryHeap.New(comparator)
	pqueue := PriorityQueue.New[int](binaryHeap)

	pqueue.Enqueue(3)                      // [3]
	pqueue.Enqueue(1,5)                    // [1,3,5]
	pqueue.Front()                     	   // 1
	pqueue.Dequeue()                       // 1
	pqueue.Length()                        // 2
	pqueue.Contains(3)                     // true
	pqueue.Contains(1)                     // false
	pqueue.Enqueue(3)                      // [1,3,5]
	pqueue.Update(1,6)                     // [3,5,6]
	pqueue.Clear()                  	   // []
	pqueue.IsEmpty()                  	   // true
	pqueue.Dequeue()                  	   // heap is empty
}
```

### Stack
A [Stack](https://en.wikipedia.org/wiki/Stack_(abstract_data_type)) is data structure that maintains a sequence of last-in, first out (LIFO). Typical operations are `push`, `pop`, and `peek`

```go
type Stack[T comparable] struct {
	_list list.List[T]
}
```

#### Generic Stack
```go
package main

import (
	"github.com/oyal2/Go-Structures/list/ArrayList"
	"github.com/oyal2/Go-Structures/stack/Stack"
)

func main() {
	stack := Stack.New[int](ArrayList.New[int](10))
    stack.Push(1)                       // 1
    stack.Push(2)                       // 2,1
    _, _ = stack.Peek()                 // 1,nil
    _, _ = stack.Pop()                  // 1,nil
    _, _ = stack.Pop()                  // 2,nil
    _, _ = stack.Pop()                  // nil,index is out of bounds
    stack.Enqueue(1)                    // 1
    stack.Push()                        // empty
    stack.IsEmpty()                     // true
    _ = stack.Length()                  // 0
}
```

### Probabilistic

#### Bloom Filter

A [Bloom Filter](https://en.wikipedia.org/wiki/Bloom_filter) is a probabilistic data structure that is used to test membership of an element in a set. It allows for false positives, but not false negatives.

```go
type BloomFilter[T utils.Ordered] struct {
	hashFuncs []func(element T, size uint64) uint64
	bits      *big.Int
	m         uint64
	size      uint64
}
```

```go
package main

import (
	"math/big"
	"github.com/oyal2/Go-Structures/BloomFilter"
)

func hashFunc1(element int, size uint64) uint64 {
	return uint64(element)
}

func hashFunc2(element int, size uint64) uint64 {
	return uint64(element) * 2
}

func main() {
	bf := BloomFilter.New[int]([]func(element int, size uint64) uint64{hashFunc1, hashFunc2}, 100)
	bf.Add(1)
	bf.Add(2)
	bf.Contains(1)        // true
	bf.Contains(2)        // true
	bf.Contains(3)        // false
	bf.Size()             // 2
	bf.Clear()            // empty
	bf.IsEmpty()          // true
	bf.Length()           // 0
	bf2 := BloomFilter.NewWithFalsePositiveRate[int]([]func(element int, size uint64) uint64{hashFunc1, hashFunc2}, 0.01, 1000)
	bf2.Add(3)
	bf2.Contains(3)       // true
	bf2.Contains(4)       // false
	bf2.Size()            // 1
}
```


### Heap

```go
type Heap[T utils.Ordered] interface {
	utils.Sort[T]
	Insert(elements ...T) error
	Extract() (T, error)
	Peek() (T, error)
	ChangeKey(element T, newElement T) error
	Contains(element T) bool
	Length() int
	IsEmpty() bool
	Clear()
}
```

#### Binary Heap

A [Binary Heap](https://en.wikipedia.org/wiki/Binary_heap) is a heap data structure that maintains the form of a binary tree. Typical operations are `Insert`, `Extract`, and `Decrease Key`. This binary heap implementation uses an array as the storage data structure. 

```go
type BinaryHeap[T utils.Ordered] struct {
	_storage    []T
	_size       int
	_comparator func(a, b T) bool
}
```

```go
package main

import "github.com/oyal2/Go-Structures/heap/BinaryHeap"

func main() {
	comparator := func(a, b int) bool {
		return a < b
	}
	binaryHeap := BinaryHeap.New(comparator)

	binaryHeap.Insert(3)                       // [3]
	binaryHeap.Insert(1,5)                     // [1,3,5]
	binaryHeap.Peek()                     	   // 1
	binaryHeap.Extract()                       // 1
	binaryHeap.Length()                        // 2
	binaryHeap.Contains(3)                     // true
	binaryHeap.Contains(1)                     // false
	binaryHeap.Insert(3)                       // [1,3,5]
	binaryHeap.ChangeKey(1,6)                  // [3,5,6]
	binaryHeap.Clear()                  	   // []
	binaryHeap.IsEmpty()                  	   // true
	binaryHeap.Extract()                  	   // heap is empty
}
```

#### Fibonacci Heap

A [Fibonacci Heap](https://en.wikipedia.org/wiki/Fibonacci_heap) is a heap data structure that maintains a collection of trees, called a heap forest. It has faster insert and merge operations compared to a binary heap, and is called a Fibonacci heap because the number of children each node has is always a Fibonacci number. Typical operations are `Insert`, `Extract`, and `Decrease Key`.

```go
type FibonacciHeap[T utils.Ordered] struct {
	_min        *Node.Node[T]
	_size       int
	_comparator func(a, b T) bool
}
```

```go
package main

import "github.com/oyal2/Go-Structures/heap/FibonacciHeap"

func main() {
	comparator := func(a, b int) bool {
		return a < b
	}
	fibHeap := FibonacciHeap.New(comparator)

	fibHeap.Insert(3)                     // [3]
	fibHeap.Insert(1,5)                   // [1,3,5]
	fibHeap.Peek()                     	  // 1
	fibHeap.Extract()                     // 1
	fibHeap.Length()                      // 2
	fibHeap.Contains(3)                   // true
	fibHeap.Contains(1)                   // false
	fibHeap.Insert(3)                     // [1,3,5]
	fibHeap.ChangeKey(1,6)                // not applicable to Fibonacci heap
	fibHeap.Clear()                  	  // []
	fibHeap.IsEmpty()                  	  // true
	fibHeap.Extract()                  	  // heap is empty
}
```
