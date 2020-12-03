// Package space implements planetary age conversion relative to Earth years.
package space

type Planet string

var planetEarthYearConversions map[Planet]float64

const EARTH_YEAR_IN_SECONDS = 31557600

// TestAge returns planetary ages in Earth years when provided with an Earth age
// in seconds.
func Age(age float64, planet Planet) float64 {
	Mercury := Planet("Mercury")
	Venus := Planet("Venus")
	Earth := Planet("Earth")
	Mars := Planet("Mars")
	Jupiter := Planet("Jupiter")
	Saturn := Planet("Saturn")
	Uranus := Planet("Uranus")
	Neptune := Planet("Neptune")

	planetEarthYearConversions := make(map[Planet]float64)
	planetEarthYearConversions[Mercury] = 0.2408467
	planetEarthYearConversions[Venus] = 0.61519726
	planetEarthYearConversions[Earth] = 1.0
	planetEarthYearConversions[Mars] = 1.8808158
	planetEarthYearConversions[Jupiter] = 11.862615
	planetEarthYearConversions[Saturn] = 29.447498
	planetEarthYearConversions[Uranus] = 84.016846
	planetEarthYearConversions[Neptune] = 164.79132

	return SecondsToEarthYears(age) / planetEarthYearConversions[planet]
}

func SecondsToEarthYears(age float64) float64 {
	return age / EARTH_YEAR_IN_SECONDS
}
