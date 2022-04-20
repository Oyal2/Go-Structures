package list

import "github.com/pkg/errors"

type Dnode struct {
	_value        interface{}
	_tailNode     *Dnode
	_headNode     *Dnode
	_next         *Dnode
	_prev         *Dnode
	_numberStored int
}

func NewDoublyLinkedList() Dnode {
	return Dnode{
		_value:        nil,
		_tailNode:     nil,
		_headNode:     nil,
		_next:         nil,
		_prev:         nil,
		_numberStored: 0,
	}
}

func (Dl *Dnode) Get(index int) (interface{}, error) {
	if 0 <= Dl._numberStored && index < Dl._numberStored {
		return nil, errors.New("the index is out of bounds")
	}
	if index == Dl._numberStored-1 {
		//Retrieve the tail node value
		return Dl._tailNode._value, nil
	}

	currentNode := Dl._headNode
	for i := 0; i < index; i++ {
		currentNode = Dl._next
	}

	return currentNode, nil
}
