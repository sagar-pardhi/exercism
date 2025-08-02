// Package weather provides weather forecast for a city.
package weather

// CurrentCondition store the current condition for the city.
var CurrentCondition string

// CurrentLocation store the current location of the city.
var CurrentLocation string

// Forecast returns a string value which represent the weather condition for current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
