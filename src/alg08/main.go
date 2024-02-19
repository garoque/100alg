package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(calculateVolumeSphere(2.5))
}

func calculateVolumeSphere(radius float64) float64 {
	if radius < 0 {
		radius *= -1
	}

	volume := (4 * math.Pi * math.Pow(radius, 3)) / 3
	return toFixed(volume, 2)
}

func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}