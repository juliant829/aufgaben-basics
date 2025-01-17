package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.
func DrawRectangle(height, width int, inner, outer string) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Printf("#")
		}
		{
			fmt.Println()
		}
		fmt.Print("B")
	}
	fmt.Println()
}

// REMARKS
// - Wenn Sie diese Aufgabe gelöst haben, können Sie die Aufgaben
//   für das Zeichnen von leeren und gefüllten Rechtecken sehr viel einfacher lösen.
