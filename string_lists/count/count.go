package count

// Count erwartet eine Liste von Strings sowie einen String, der gezählt werden soll.
// Es liefert die Anzahl der Vorkommen des gesuchten Strings in der Liste.
func Count(strings []string, search string) int {
	count := 0

	for _, s := range strings {
		if s == search {
			count++
		}
	}
	return count
}
