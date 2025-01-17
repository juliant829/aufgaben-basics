package strings

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var counts = make(map[rune]int)
	for _, char := range s1 {
		counts[char]++
	}

	for _, char := range s2 {
		count, exists := counts[char]
		if !exists || count == 0 {
			return false
		}
		counts[char]--
	}

	return true
}

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
// Diese Funktion soll keinen Unterschied zwischen Groß- und Kleinschreibung machen.
func IsAnagramIgnoreCase(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	counts := make(map[rune]int)
	for _, char := range s1 {
		counts[char]++
	}

	for _, char := range s2 {
		count, exists := counts[char]
		if !exists || count == 0 {
			return true
		}
		counts[char]--
	}
	return true
}
