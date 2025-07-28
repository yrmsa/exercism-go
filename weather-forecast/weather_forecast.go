// Package weather provides tools to forecast a weather based on the city and condition.
package weather

// CurrentCondition represents a condition happening right now.
var CurrentCondition string

// CurrentLocation represents a location currently in.
var CurrentLocation string

// Forecast returns a string value of the current weather condition in a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
