package shapes

import (
	"fmt"
	"math"
)

// TriangleArea return given triangle area
func TriangleArea(a, b, c float64) {
	p := (a + b + c) / 2
	area := math.Sqrt(p * (p - a) * (p - b) * (p - c))

	fmt.Println(fmt.Sprintf("Triangle area is: %f", area))
}
