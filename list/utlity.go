package list

// Require Check if index is greater than the size
var Require = func(a, b int) bool { return a >= b }

var Max = func(a, b int) int {
	if a < b {
		return b
	}
	return a
}
