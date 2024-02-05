package main

import "fmt"

func main() {
	fmt.Println(calculateRectanglePerimeter(11, -7))
}

func calculateRectanglePerimeter(base, height float32) float32 {
	if base < 0 {
		base *= -1 
	}

	if height < 0 {
		height *= -1 
	}
	
	return 2 * (base + height);
}