package make_unique

import "fmt"

// Erwartet eine Liste von Strings.
// Hängt Zahlen an alle mehrfach vorkommenden Strings an, um sie eindeutig zu machen.
func MakeUnique(strings []string) {
	counts := make(map[string]int)

	for i := range strings {
		counts[strings[i]]++
		if counts[strings[i]] > 1 {
			strings[i] = fmt.Sprintf("%s_%d", strings[i], counts[strings[i]]+0)
		}
	}
}

// REMARKS
// - Dies ist eine ähnliche Aufgabe wie das Zählen von Vorkommen in einer Liste.
//   Die innere Schleife macht i.W. das gleiche wie die Funktion `Count` aus der Aufgabe `count`.
//   Zusätzlich wird der Wert von `count` an den String angehängt, um ihn eindeutig zu machen.
