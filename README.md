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
  - [Lists](#lists)
  - [Tree](#tree)
  - [Hashed Data Structures](#hashed-data-structures)
  - [Heap](#heap)
  - [Queue](#queue)
  - [Stack](#stack)
- [🔧 Benchmarks ](#-benchmarks-)
  - [ArrayList](#arraylist)
  - [SinglyLinked List](#singlylinked-list)
- [🎈 Usage ](#-usage-)
  - [List Interface](#list-interface)
    - [ArrayList](#arraylist-1)
    - [Fixed ArrayList](#fixed-arraylist)
  - [Node Type](#node-type)
    - [Singly Linked List](#singly-linked-list)
    - [Doubly Linked List](#doubly-linked-list)
  - [Queue](#queue-1)
    - [Generic Queue](#generic-queue)
  - [Stack](#stack-1)
    - [Generic Stack](#generic-stack)

## 🧐 About <a name = "about"></a>
I made this library because Golang has a very limited amount of data structures. I implemented various data structures such as ArrayList, Queue, DoublyLinked List, SinglyLinked List, and many more. There are many complex data structures in this library and some basic ones. All Golang data structures are optimized and are generic classes. 

## ⚙️ Data Structures <a name = "data_structures"></a>
### Lists

- [x] ArrayList
- [x] Singly Linked List
- [x] Doubly Linked List
- [x] Fixed Array List

### Tree
- [x] AVL Tree
- [ ] B-Tree
- [ ] Trei

### Hashed Data Structures
- [x] Bloom Filter
- [ ] Hash Table

### Heap
- [ ] Heap
- [ ] Fibonacci Heap

### Queue
- [x] Queue
- [ ] Circular Queue
- [ ] Priority Queue

### Stack
- [x] Stack

## 🔧 Benchmarks <a name = "benchmarks"></a>

Benchmark Specs 
```
goos: windows
goarch: amd64
cpu: AMD Ryzen 9 5950X 16-Core Processor     
```
### ArrayList
emirpasic/gods ArrayList:
```      
BenchmarkArrayListGet100          	21818300	        55.73 ns/op
BenchmarkArrayListGet1000        	 2539680	       460.1 ns/op
BenchmarkArrayListGet10000       	  266656	      4575 ns/op
BenchmarkArrayListGet100000      	   25882	     45321 ns/op
BenchmarkArrayListAdd100         	  489795	      2640 ns/op
BenchmarkArrayListAdd1000        	   49704	     33146 ns/op
BenchmarkArrayListAdd10000       	    3428	    429405 ns/op
BenchmarkArrayListAdd100000       	     422	   3704970 ns/op
BenchmarkArrayListRemove100      	44444937	        28.30 ns/op
BenchmarkArrayListRemove1000      	 5106406	       231.0 ns/op
BenchmarkArrayListRemove10000    	  524035	      2313 ns/op
BenchmarkArrayListRemove100000    	       1	25059500300 ns/op
```

GoStructure ArrayList:
```       
BenchmarkArrayListGet100          	42104228	        27.44 ns/op
BenchmarkArrayListGet1000         	 5172422	       234.1 ns/op
BenchmarkArrayListGet10000        	  521720	      2256 ns/op
BenchmarkArrayListGet100000       	   53810	     22644 ns/op
BenchmarkArrayListAdd100          	  694474	      2167 ns/op
BenchmarkArrayListAdd1000         	   35397	     30313 ns/op
BenchmarkArrayListAdd10000        	    3927	    262669 ns/op
BenchmarkArrayListAdd100000       	     378	   2982804 ns/op
BenchmarkArrayListRemove100       	  118226	     10057 ns/op
BenchmarkArrayListRemove1000      	   10000	    114950 ns/op
BenchmarkArrayListRemove10000     	    1029	   1180761 ns/op
BenchmarkArrayListRemove100000    	      99	  12146458 ns/op
```

![image](https://user-images.githubusercontent.com/13637813/195453768-c6f0638c-cfff-412f-9dfe-2f406d6280c3.png)
![image](https://user-images.githubusercontent.com/13637813/195451740-8276f97e-f96b-40bc-abb7-104bb378b79a.png)

### SinglyLinked List
emirpasic/gods SinglyLinked List:
```      
BenchmarkSinglyLinkedListGet100          	  436363	      2602 ns/op
BenchmarkSinglyLinkedListGet1000         	    2757	    424012 ns/op
BenchmarkSinglyLinkedListGet10000        	      25	  44679900 ns/op
BenchmarkSinglyLinkedListGet100000       	       1	4524499900 ns/op
BenchmarkSinglyLinkedListAdd100          	  196718	      6120 ns/op
BenchmarkSinglyLinkedListAdd1000         	   21997	     56667 ns/op
BenchmarkSinglyLinkedListAdd10000        	    2221	    591400 ns/op
BenchmarkSinglyLinkedListAdd100000       	     212	   5780662 ns/op
BenchmarkSinglyLinkedListRemove100       	 8856133	       135.4 ns/op
BenchmarkSinglyLinkedListRemove1000      	  888882	      1313 ns/op
BenchmarkSinglyLinkedListRemove10000     	   89887	     13144 ns/op
BenchmarkSinglyLinkedListRemove100000    	    8888	    130794 ns/op
```

GoStructure SinglyLinked List:
```       
BenchmarkSinglyLinkedListGet100          	  444448	      2565 ns/op
BenchmarkSinglyLinkedListGet1000         	    2962	    409858 ns/op
BenchmarkSinglyLinkedListGet10000        	      25	  44960028 ns/op
BenchmarkSinglyLinkedListGet100000       	       1	4552001100 ns/op
BenchmarkSinglyLinkedListAdd100          	  200001	      5785 ns/op
BenchmarkSinglyLinkedListAdd1000         	   15675	     64625 ns/op
BenchmarkSinglyLinkedListAdd10000        	    1882	    849097 ns/op
BenchmarkSinglyLinkedListAdd100000       	     156	   6519228 ns/op
BenchmarkSinglyLinkedListRemove100       	  615418	      1995 ns/op
BenchmarkSinglyLinkedListRemove1000      	   61380	     19916 ns/op
BenchmarkSinglyLinkedListRemove10000     	    5454	    205904 ns/op
BenchmarkSinglyLinkedListRemove100000    	     446	   2540357 ns/op
```

## 🎈 Usage <a name = "usage"></a>

### List Interface

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
A [Queue](https://en.wikipedia.org/wiki/Queue_(abstract_data_type)) is data structure that maintains a sequence of first-in-first-out (FIFO). Typical operations are `enqueue`, `dequeue`, and `front`

```go
type Queue[T comparable] struct {
	_list list.List[T]
}
```

#### Generic Queue
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