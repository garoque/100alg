package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(calculateAreaSquare(-11))
}

func calculateAreaSquare(l float64) float64 {
	return math.Pow(l, 2)
}