// Package weather provides useful information for forecasts weather.
package weather

// CurrentCondition describes current condition.
var CurrentCondition string

// CurrentLocation describes current location.
var CurrentLocation string

// Forecasts returns the forecast for the city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
