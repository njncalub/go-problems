// Package space calculates how old someone would be on Earth, Mercury, Venus,
// Mars, Jupiter, Saturn, Uranus, and Neptune.
package space

type Planet string

// Define multipliers for different Planets.
var (
	EarthYear   float64 = 31557600
	MercuryYear float64 = 0.2408467 * EarthYear
	VenusYear   float64 = 0.61519726 * EarthYear
	MarsYear    float64 = 1.8808158 * EarthYear
	JupiterYear float64 = 11.862615 * EarthYear
	SaturnYear  float64 = 29.447498 * EarthYear
	UranusYear  float64 = 84.016846 * EarthYear
	NeptuneYear float64 = 164.79132 * EarthYear
)

// Age returns your actual age on a specific Planet.
func Age(seconds float64, planet Planet) (age float64) {
	switch planet {
	case "Earth":
		age = seconds / EarthYear
	case "Mercury":
		age = seconds / MercuryYear
	case "Venus":
		age = seconds / VenusYear
	case "Mars":
		age = seconds / MarsYear
	case "Jupiter":
		age = seconds / JupiterYear
	case "Saturn":
		age = seconds / SaturnYear
	case "Uranus":
		age = seconds / UranusYear
	case "Neptune":
		age = seconds / NeptuneYear
	}
	return age
}
