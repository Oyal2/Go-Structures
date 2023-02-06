package list

type List[T comparable] interface {
	Get(index int) (T, error)
	Insert(index int, element T) error
	Remove(index int) (T, error)
	Length() int
	Update(index int, element T) (T, error)
	IndexOf(element T) int
	Clear()
}

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}
