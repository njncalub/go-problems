// Package space calculates how old someone would be on Earth, Mercury, Venus,
// Mars, Jupiter, Saturn, Uranus, and Neptune.
package space

type Planet string

// Define total number of seconds in one Earth year.
const EarthYear = 31557600

// Define multipliers for different Planets.
var Seconds = map[Planet]float64{
	"Earth":   EarthYear * 1.0,
	"Mercury": EarthYear * 0.2408467,
	"Venus":   EarthYear * 0.61519726,
	"Mars":    EarthYear * 1.8808158,
	"Jupiter": EarthYear * 11.862615,
	"Saturn":  EarthYear * 29.447498,
	"Uranus":  EarthYear * 84.016846,
	"Neptune": EarthYear * 164.79132,
}

// Age returns someone's actual age on a specific Planet.
func Age(seconds float64, planet Planet) float64 {
	return seconds / Seconds[planet]
}
