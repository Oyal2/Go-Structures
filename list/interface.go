package list

type List interface {
	Get(index int) (interface{}, error)
	Insert(index int, element interface{}) error
	Remove(index int) (interface{}, error)
	Length() int
}
