package utils

type Sort[T any] interface {
	Less(a, b T) bool
}

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
