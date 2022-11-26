package utils

type Ordered interface {
	~int | ~uint | ~int8 | ~uint8 | ~int16 | ~uint16 |
		~int32 | ~uint32 | ~int64 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}
