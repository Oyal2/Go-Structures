package FibonacciHeap

import (
	"errors"
	"fmt"
	"math"
	"sync"

	"github.com/oyal2/Go-Structures/heap/FibonacciHeap/Node"
	"github.com/oyal2/Go-Structures/utils"
)

type FibonacciHeap[T utils.Ordered] struct {
	_min        *Node.Node[T]
	_size       int
	_comparator func(a, b T) bool
	sync.RWMutex
}

func New[T utils.Ordered](comparator func(a, b T) bool) *FibonacciHeap[T] {
	return &FibonacciHeap[T]{
		_min:        nil,
		_size:       0,
		_comparator: comparator,
	}
}

func (FH *FibonacciHeap[T]) Insert(elements ...T) error {
	FH.Lock()
	defer FH.Unlock()
	for _, element := range elements {
		err := FH.insert_item(&Node.Node[T]{
			Element: element,
		})
		if err != nil {
			return err
		}
		FH._size++
	}
	return nil
}

func (FH *FibonacciHeap[T]) insert_item(node *Node.Node[T]) error {
	node.Degree = 0
	node.Parent = nil
	if FH._min == nil {
		FH._min = node
		node.Left = FH._min
		node.Right = FH._min
	} else {
		node.Right = FH._min
		node.Left = FH._min.Left
		FH._min.Left.Right = node
		FH._min.Left = node
		// update the minimum node if necessary
		if FH._comparator(node.Element, FH._min.Element) {
			FH._min = node
		}
	}

	return nil
}

func (FH *FibonacciHeap[T]) Extract() (returnValue T, err error) {
	FH.Lock()
	defer FH.Unlock()
	if FH._size == 0 {
		return returnValue, errors.New("empty heap")
	}
	z := FH._min

	if z != nil {
		childNode := z.Child
		k := childNode
		if childNode != nil {
			for {
				childRight := childNode.Right
				FH.insert_item(childNode)
				childNode = childRight
				if childNode == nil || childNode == k {
					break
				}
			}
		}
		z.Left.Right = z.Right
		z.Right.Left = z.Left
		z.Child = nil
		if z == z.Right {
			FH._min = nil
		} else {
			FH._min = z.Right
			FH.consolidate()
		}
		FH._size--
		returnValue = z.Element

	}
	return returnValue, err
}

func (FH *FibonacciHeap[T]) consolidate() {
	nodeArr := make([]*Node.Node[T], FH.maxDegree())

	w := FH._min
	last := w.Left
	for {
		rightW := w.Right
		x := w
		d := x.Degree
		for nodeArr[d] != nil {
			y := nodeArr[d]
			if !FH._comparator(x.Element, y.Element) {
				x, y = y, x
			}
			FH.heapLink(y, x)
			nodeArr[d] = nil
			d++
		}
		nodeArr[d] = x
		if w == last {
			break
		}
		w = rightW
	}
	FH._min = nil

	for i := 0; i < len(nodeArr); i++ {
		if nodeArr[i] != nil {
			FH.insert_item(nodeArr[i])
		}
	}
}

func (FH *FibonacciHeap[T]) heapLink(yChild, xParent *Node.Node[T]) {
	yChild.Left.Right = yChild.Right
	yChild.Right.Left = yChild.Left

	if xParent.Child == nil {
		xParent.Child = yChild
		yChild.Right = yChild
		yChild.Left = yChild
	} else {
		xParent.Child.Left.Right = yChild
		yChild.Right = xParent.Child
		yChild.Left = xParent.Child.Left
		xParent.Child.Left = yChild
	}
	yChild.Parent = xParent
	xParent.Child = yChild
	xParent.Degree++
	yChild.Mark = false
}

func (FH *FibonacciHeap[T]) maxDegree() int {
	phi := (1 + math.Sqrt(5)) / 2
	return int(math.Floor(math.Log(float64(FH._size)) / math.Log(phi)))
}

func (FH *FibonacciHeap[T]) Contains(element T) bool {
	FH.RLock()
	defer FH.RUnlock()
	if FH._size == 0 {
		return false
	}

	return FH.find(element, FH._min) != nil
}

func (FH *FibonacciHeap[T]) ChangeKey(element T, newElement T) error {
	FH.Lock()
	defer FH.Unlock()
	if FH._size == 0 {
		return errors.New("empty heap")
	}

	node := FH.find(element, FH._min)
	if node == nil {
		return errors.New("element doesnt exist")
	}

	if !FH._comparator(newElement, element) {
		return errors.New("not applicable to heap")
	}

	node.Element = newElement
	nodeY := node.Parent

	if nodeY != nil && FH._comparator(node.Element, nodeY.Element) {
		FH.cut(node, nodeY)
		FH.cascadingCut(nodeY)
	}

	if FH._comparator(node.Element, FH._min.Element) {
		FH._min = node
	}

	return nil
}

func (FH *FibonacciHeap[T]) find(element T, node *Node.Node[T]) *Node.Node[T] {
	if node == nil {
		return nil
	}

	currNode := node
	for {
		foundNode := FH.find(element, currNode.Child)

		if foundNode != nil {
			return foundNode
		}

		if currNode.Element == element {
			return currNode
		}
		if node == currNode.Right {
			break
		}
		currNode = currNode.Right
	}

	return nil
}

func (FH *FibonacciHeap[T]) cut(nodeX, nodeY *Node.Node[T]) {
	nodeX.Right.Left = nodeX.Left
	nodeX.Left.Right = nodeX.Right
	nodeY.Degree--
	nodeX.Right, nodeX.Left = nil, nil
	FH.insert_item(nodeX)
	nodeX.Mark = false
}

func (FH *FibonacciHeap[T]) cascadingCut(nodeY *Node.Node[T]) {
	z := nodeY.Parent

	if z != nil {
		if !nodeY.Mark {
			nodeY.Mark = true
		} else {
			FH.cut(z, nodeY)
			FH.cascadingCut(z)
		}
	}
}

func (FH *FibonacciHeap[T]) Peek() (returnValue T, err error) {
	FH.Lock()
	defer FH.Unlock()
	if FH._size == 0 {
		return returnValue, errors.New("heap is empty")
	}

	returnValue = FH._min.Element
	return returnValue, nil
}

func (FH *FibonacciHeap[T]) IsEmpty() bool {
	return FH._size == 0
}

func (FH *FibonacciHeap[T]) Less(a, b T) bool {
	return FH._comparator(a, b)
}

func (FH *FibonacciHeap[T]) Length() int {
	FH.RLock()
	defer FH.RUnlock()
	return FH._size
}

func (FH *FibonacciHeap[T]) Clear() {
	FH.Lock()
	defer FH.Unlock()
	FH._min = nil
	FH._size = 0
}

func (FH *FibonacciHeap[T]) Display() {
	FH.RLock()
	defer FH.RUnlock()
	FH.traverseDisplay(FH._min, 0)
	fmt.Println()
}

func (FH *FibonacciHeap[T]) traverseDisplay(node *Node.Node[T], depth int) {
	if node == nil {
		return
	}
	var items []T
	currNode := node
	for {
		FH.traverseDisplay(currNode.Child, depth+1)
		items = append(items, currNode.Element)
		if node == currNode.Right {
			break
		}
		currNode = currNode.Right
	}
	fmt.Printf("Depth %d: ", depth)

	fmt.Print(items)
	fmt.Println()
}
