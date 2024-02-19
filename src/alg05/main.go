package main

import "fmt"

func main() {
	fmt.Println(calculateTrianguleArea(11, -7))
}

func calculateTrianguleArea(base, height float32) float32 {
	if base < 0 {
		base *= -1 
	}

	if height < 0 {
		height *= -1 
	}
	
	return (base * height) / 2;
}