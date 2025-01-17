package forms

import "fmt"

// ExampleAreaSquare demonstrates how to calculate the area of a square
// given the length of its side 'a'. The formula used is a^2.
func ExampleAreaSquare(a float64) {
	fmt.Println(AreaSquare(0))
	fmt.Println(AreaSquare(1))
	fmt.Println(AreaSquare(2))
	fmt.Println(AreaSquare(3))
	fmt.Println(AreaSquare(4))
	fmt.Println(AreaSquare(5))

	// Output:
	// 0
	// 1
	// 4
	// 9
	// 16
	// 25
}
