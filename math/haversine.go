package math

// Implements Haversine formula
// https://www.math.ksu.edu/~dbski/writings/haversine.pdf

import (
    "github.com/dantheman213/go-gps/location"
    "math"
)

const (
    earthRadiusKM = 6371
)

func DegreesToRadians(d float64) float64 {
    return d * math.Pi / 180
}

func CalculateDistanceBetweenPointsInKM(c1, c2 *location.LocationDD) float64 {
    lat1 := DegreesToRadians(c1.Latitude)
    lon1 := DegreesToRadians(c1.Longitude)
    lat2 := DegreesToRadians(c2.Latitude)
    lon2 := DegreesToRadians(c2.Longitude)

    diffLat := lat2 - lat1
    diffLon := lon2 - lon1

    a := math.Pow(math.Sin(diffLat / 2), 2) + math.Cos(lat1)*math.Cos(lat2) * math.Pow(math.Sin(diffLon / 2), 2)
    c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1 - a))
    distanceKM := c * earthRadiusKM

    return distanceKM
}
