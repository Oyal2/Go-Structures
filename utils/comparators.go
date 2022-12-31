package utils

type Sort[T Ordered] interface {
	Less(a, b T) bool
	Swap(i, j any)
}

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
