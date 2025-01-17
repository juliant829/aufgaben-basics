package triangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Das Dreieck soll komplett mit `#`-Zeichen gefüllt sein.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
func DrawSolidTriangle(length int) {
	for i := 0; i < length; i++ {
		fmt.Print("#")
		for j := 0; j < i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
	fmt.Println()
}
