package main

import "fmt"

func main() {
	fmt.Println(conventMetersToCentimeters(11))
}

func conventMetersToCentimeters(meter float32) float32 {
	if meter < 0 {
		meter *= -1 
	}

	return meter * 100;
}