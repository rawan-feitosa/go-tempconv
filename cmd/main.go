package main

import "fmt"

// Offset to convert kelvins to celsius / celsius to kelvins
var kelvinOffset float64 = 273.15

// function to convert kelvins to celsius
func KelvinToCelsius(value float64) float64 {
	return value - kelvinOffset
}

// function to convert celsius to kelvins
func CelsiusToKelvin(value float64) float64 {
	return value + kelvinOffset
}

func main() {
	var actualTemp float64 = 30 // In Celsius
	fmt.Printf("Original Value in: %v\n", actualTemp)

	kelvinValue := CelsiusToKelvin(actualTemp)
	fmt.Printf("ACTUAL TEMPERATURE IN KELVIN: %v\n\n", kelvinValue)

	celsiusValue := KelvinToCelsius(kelvinValue)
	fmt.Printf("ACTUAL TEMPERATURE IN CELSIUS: %v\n", celsiusValue)
}
