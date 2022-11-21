package list

type List interface {
	Get(index int) (interface{}, error)
	Insert(index int, element interface{}) error
	Remove(index int) (interface{}, error)
	Length() int
}

type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}
