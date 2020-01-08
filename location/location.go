package location

// DD (Decimal Degrees)
type LocationDD struct {
    Latitude  float64
    Longitude float64
}

// DDM (Degrees Decimal Minutes)
type LocationDDM struct {
    Latitude           float64
    LatitudeDirection  string
    Longitude          float64
    LongitudeDirection string
}
