package utils

func InRange(min int, max int, x int) bool {
	if x <= max && x >= min {
		return true
	} else {
		return false
	}
}
