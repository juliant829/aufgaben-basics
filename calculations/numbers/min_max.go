package numbers

// Berechnet das Minimum von zwei Zahlen.
func Min(a, b int) int {
	// TODO
	if a < b {
		return a
	}
	if a > b {
		return b
	}
	if a == b {
		return b
	}
	return 0
}

// Berechnet das Maximum von zwei Zahlen.
func Max(a, b int) int {
	// TODO
	if a > b {
		return a
	}
	if a < b {
		return b
	}
	if a == b {
		return b
	}
	return 0
}
