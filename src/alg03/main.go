package main

import "fmt"

func main() {
	fmt.Println(convertCelsiusToFahrenheit(32))
}

func convertCelsiusToFahrenheit(celsius float32) float32 {
	return (celsius * 9) / 5 + 32;
}