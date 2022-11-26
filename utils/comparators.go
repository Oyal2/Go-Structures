package utils

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
