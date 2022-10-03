package list

import "errors"

type Snode struct {
	_value        interface{}
	_tailNode     *Snode
	_headNode     *Snode
	_next         *Snode
	_numberStored *int
}

func NewSinglyLinkedList() Snode {
	i := 0
	return Snode{
		_value:        nil,
		_tailNode:     nil,
		_headNode:     nil,
		_next:         nil,
		_numberStored: &i,
	}
}


func (Sl *Snode) Get(index int) (interface{}, error) {
	if 0 > index || index >= *Sl._numberStored {
		return nil, errors.New("the index is out of bounds")
	}
	if index == *Sl._numberStored-1 {
		//Retrieve the tail node value
		return Sl._tailNode._value, nil
	}

	currentNode := Sl._headNode
	for i := 0; i < index; i++ {
		currentNode = currentNode._next
	}

	return currentNode._value, nil
}

func (Sl *Snode) Update(index int, element interface{}) error {
	if 0 > index || index >= *Sl._numberStored {
		return errors.New("the index is out of bounds")
	}
	if index == *Sl._numberStored-1 {
		Sl._tailNode._value = element
		return nil
	}

	if index == 0 {
		Sl._value = element
		return nil
	}

	currentNode := Sl._headNode
	for i := 0; i < index; i++ {
		currentNode = currentNode._next
	}
	currentNode._value = element
	return nil
}

func (Sl *Snode) Insert(index int, element interface{}) error {
	if 0 > index || index > *Sl._numberStored {
		return errors.New("the index is out of bounds")
	}
	newNode := NewSinglyLinkedList()
	newNode._value = element
	//If the SL is empty
	if *Sl._numberStored == 0 {
		Sl._headNode = Sl
		Sl._tailNode = Sl
		Sl._value = element
		*Sl._numberStored++
		newNode._numberStored = Sl._numberStored
		return nil
	} else {
		if index == 0 {
			newNode._next = Sl._headNode
			Sl._headNode = &newNode
		} else if index < *Sl._numberStored {
			currentNode := Sl._headNode
			for i := 0; i < (index - 1); i++ {
				currentNode = currentNode._next
			}
			newNode._next = currentNode._next
			currentNode._next = &newNode
		} else {
			//Insert at tail
			newNode._next = nil
			newNode._headNode = Sl._headNode
			Sl._tailNode._next = &newNode
			Sl._tailNode = Sl._tailNode._next
			if 1 == Sl.Length() {
				Sl._next = Sl._tailNode
			}
		}
		*Sl._numberStored++
		newNode._numberStored = Sl._numberStored
	}
	return nil
}

func (Sl *Snode) Remove(index int) (interface{}, error) {
	if 0 > index || index >= *Sl._numberStored {
		return nil, errors.New("the index is out of bounds")
	}

	if *Sl._numberStored == 1 {
		result := Sl._headNode._value
		Sl._headNode = nil
		Sl._tailNode = nil
		*Sl._numberStored--
		return result, nil
	} else {
		if index == 0 {
			result := Sl._headNode._value
			Sl._headNode = Sl._headNode._next
			*Sl._numberStored--
			return result, nil
		} else if index == *Sl._numberStored-1 {
			result := Sl._tailNode._value
			currentNode := Sl._headNode
			for i := 0; i < (index - 1); i++ {
				currentNode = currentNode._next
			}
			Sl._tailNode = currentNode
			Sl._tailNode._next = nil
			*Sl._numberStored--
			return result, nil
		} else {
			currentNode := Sl._headNode
			for i := 0; i < (index - 1); i++ {
				currentNode = currentNode._next
			}
			result := currentNode._next._value
			currentNode._next = currentNode._next._next
			*Sl._numberStored--
			return result, nil
		}
	}
}

func (Sl *Snode) Length() int { return *Sl._numberStored }