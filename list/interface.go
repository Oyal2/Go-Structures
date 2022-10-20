package list

type List interface {
	Get(index int) (interface{}, error)
	Insert(index int, element interface{}) error
	Remove(index int) (interface{}, error)
	Length() int
}

type Node struct {
	_value interface{}
	_next  *Node
	_prev  *Node
}
