package SinglyLinkedList

import (
	"errors"

	"github.com/oyal2/Go-Structures/list"
)

type SinglyLinkedList[T comparable] struct {
	tailNode     *list.Node[T]
	headNode     *list.Node[T]
	numberStored int
}

func New[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (Sl *SinglyLinkedList[T]) Get(index int) (value T, err error) {
	if 0 > index || index >= Sl.numberStored {
		return value, errors.New("the index is out of bounds")
	}
	if index == Sl.numberStored-1 {
		//Retrieve the tail node value
		return Sl.tailNode.Value, nil
	}

	currentNode := Sl.headNode
	for i := 0; i != index; i++ {
		currentNode = currentNode.Next
	}

	return currentNode.Value, nil
}

func (Sl *SinglyLinkedList[T]) Update(index int, element T) (T, error) {
	var empty T
	if 0 > index || index >= Sl.numberStored {
		return empty, errors.New("the index is out of bounds")
	}
	if index == Sl.numberStored-1 {
		Sl.tailNode.Value = empty
		Sl.tailNode.Value = element
		return empty, nil
	}

	if index == 0 {
		Sl.headNode.Value = empty
		Sl.headNode.Value = element
		return empty, nil
	}

	currentNode := Sl.headNode
	for i := 0; i != index; i++ {
		currentNode = currentNode.Next
	}
	currentNode.Value = empty
	currentNode.Value = element
	return empty, nil
}

func (Sl *SinglyLinkedList[T]) Insert(index int, element T) error {
	if 0 > index || index > Sl.numberStored {
		return errors.New("the index is out of bounds")
	}
	newNode := list.Node[T]{Value: element}
	//If the SL is empty
	if Sl.numberStored == 0 {
		Sl.headNode = &newNode
		Sl.tailNode = &newNode
		Sl.numberStored++
		return nil
	}

	if index == 0 {
		newNode.Next = Sl.headNode
		Sl.headNode = &newNode
		Sl.numberStored++
		return nil
	}

	if index == Sl.numberStored {
		//Insert at tail
		Sl.tailNode.Next = &newNode
		Sl.tailNode = &newNode
		Sl.numberStored++
		return nil
	}

	currentNode := Sl.headNode
	for i := 0; i != index; i++ {
		currentNode = currentNode.Next
	}
	newNode.Next = currentNode.Next
	currentNode.Next = &newNode

	Sl.numberStored++
	return nil
}

func (Sl *SinglyLinkedList[T]) Add(elements ...T) error {
	return Sl.Append(elements...)
}

func (Sl *SinglyLinkedList[T]) Prepend(elements ...T) error {
	for _, element := range elements {
		err := Sl.Insert(0, element)
		if err != nil {
			return err
		}
	}
	return nil
}

func (Sl *SinglyLinkedList[T]) Append(elements ...T) error {
	for _, element := range elements {
		err := Sl.Insert(Sl.Length(), element)
		if err != nil {
			return err
		}
	}
	return nil
}

func (Sl *SinglyLinkedList[T]) Contains(element T) bool {
	currentNode := Sl.headNode
	for currentNode != nil {
		if currentNode.Value == element {
			return true
		}
		currentNode = currentNode.Next
	}
	return false
}

func (Sl *SinglyLinkedList[T]) IndexOf(element T) int {
	currentNode := Sl.headNode
	for i := 0; currentNode != nil; i++ {
		if currentNode.Value == element {
			return i
		}
		currentNode = currentNode.Next
	}
	return -1
}

func (Sl *SinglyLinkedList[T]) Clear() {
	Sl.headNode = nil
	Sl.tailNode = nil
	Sl.numberStored = 0
}

func (Sl *SinglyLinkedList[T]) Empty() bool {
	return Sl.Length() == 0
}

func (Sl *SinglyLinkedList[T]) Remove(index int) (value T, err error) {
	if 0 > index || index >= Sl.numberStored {
		return value, errors.New("the index is out of bounds")
	}

	var result T
	if Sl.numberStored == 1 {
		result = Sl.headNode.Value
		Sl.headNode = nil
		Sl.tailNode = nil
		Sl.numberStored = 0
		return result, nil
	}

	if index == 0 {
		result = Sl.headNode.Value
		Sl.headNode = Sl.headNode.Next
		Sl.numberStored--
		return result, nil
	}

	currentNode := Sl.headNode
	for e := 0; e != index-1; e, currentNode = e+1, currentNode.Next {
	}

	result = currentNode.Next.Value
	currentNode.Next = currentNode.Next.Next

	if index == Sl.numberStored-1 {
		Sl.tailNode = currentNode
		Sl.tailNode.Next = nil
	}

	Sl.numberStored--
	return result, nil

}

func (Sl *SinglyLinkedList[T]) Length() int { return Sl.numberStored }
