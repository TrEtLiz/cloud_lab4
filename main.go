package main

import (
	"fmt"
	"math"
)

// distance function to calculate the distance between two points
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}


func main() {
	// calling the distance function
	fmt.Println(distance(0, 0, 3, 4))
}