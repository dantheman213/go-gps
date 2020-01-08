package convert

// Kilometers per hour to miles per Hour
func KPHToMPH(kph float64) float64 {
    return kph / 1.609
}

// Kilometers per hour to meters per second
func KPHToMPS(kph float64) float64 {
    return kph / 3.6
}

// Kilometers per hour to knots (nautical mile per hour)
func KPHToKnots(kph float64) float64 {
    return kph / 1.852
}
