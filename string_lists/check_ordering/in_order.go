package check_ordering

// Erwartet eine Liste von Strings und zwei einzelne Strings.
// Überprüft, ob die beiden Strings in der Liste in der gegebenen Reihenfolge vorkommen.
// Gibt `true` zurück, wenn das der Fall ist, ansonsten `false`.
func CheckOrdering(strings []string, first, second string) bool {
	pos1 := -1
	pos2 := -1
	for i, s := range strings {
		if s == first {
			pos1 = i
		}
		if s == second {
			pos2 = i
		}
	}
	return pos1 != -1 && pos2 != -1 && pos1 < pos2
}

// REMARKS
// - Diese Aufgabe ist eine komplexere Variante der Aufgabe "Prüfen, ob ein Element in einer Liste vorkommt".
// - Sie können die Lösung der einfachen Variante als Grundlage verwenden und diese entsprechend erweitern.
