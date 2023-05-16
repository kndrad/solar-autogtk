// Package sunphase provides functions to calculate sunrise and sunset times.
package sunphase

import (
	"time"

	"github.com/mourner/suncalc-go"
)

// Sunrise calculates the time of sunrise for the provided latitude and longitude.
// The time is calculated for the current day.
// It returns the time of sunrise as a time.Time value.
func Sunrise(latitude, longitude float64) time.Time {
	times := suncalc.SunTimes(time.Now(), latitude, longitude)
	return times["sunrise"]
}

// Sunset calculates the time of sunset for the provided latitude and longitude.
// The time is calculated for the current day.
// It returns the time of sunset as a time.Time value.
func Sunset(latitude, longitude float64) time.Time {
	times := suncalc.SunTimes(time.Now(), latitude, longitude)
	return times["sunset"]
}
