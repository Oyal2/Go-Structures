package BTree

import (
	"testing"
)

func TestInsert1(t *testing.T) {
	list := New[int](3)
	for i := 1; i < 7; i++ {
		list.Insert(i)
	}
	return
}
