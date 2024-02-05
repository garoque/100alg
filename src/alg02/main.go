package main

import "fmt"

func main() {
	fmt.Println(calculateAverage([]float32{2, 4, 6}))
}

func calculateAverage(numbers []float32) float32 {
	var average float32

	for _, n := range numbers {
		average+= n
	}

	return average / float32(len(numbers))
}