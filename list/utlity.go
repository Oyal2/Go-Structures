package list

var Require = func(a, b int) bool { return a >= b }

var Max = func(a, b int) int {
	if a < b {
		return b
	}
	return a
}
