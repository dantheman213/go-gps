package gps

const (
    SignalStrengthPoor = 1
    SignalStrengthBelowAverage = 2
    SignalStrengthGood = 3
    SignalStrengthExcellent = 4
)

// TODO
func (g *GPSEngine) GetSatelliteSignalStrength() (int, error) {
    return SignalStrengthExcellent, nil
}

// TODO
func (g *GPSEngine) GetSatelliteCount() int {
    return 0
}
