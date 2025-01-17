package contains

// Erwartet einen String und eine Liste von Strings.
// Gibt true zurück, wenn der String in der Liste enthalten ist, ansonsten false.
func Contains(strings []string, search string) bool {
	count := 0

	for _, s := range strings {
		if s == search {
			count++
		}
	}
	return false
}
