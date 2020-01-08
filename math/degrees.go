package math

import (
    "errors"
    "math"
)

// DDM (Degrees Decimal Minutes) to DD (Decimal Degrees)
func DDMToDD(value float64, direction string) (float64, error) {
    if direction != "N" && direction != "S" && direction != "E" && direction != "W" {
        return 0, errors.New("direction is invalid")
    }

    degrees := math.Floor(value / 100)
    minutes := ((value / 100) - math.Floor(value/100)) * 100 / 60
    decimal := degrees + minutes

    if direction == "W" || direction == "S" {
        decimal *= -1
    }

    return float64(decimal), nil
}

// TODO
func DDMToDMS(value float64, direction string) (degrees, minutes, seconds float64) {

    return 0, 0, 0
}

func DDToDMS(value float64) (degrees, minutes, seconds int) {
    degrees = int(value)
    minutes = int((value - float64(degrees)) * 60)
    seconds = int((value - float64(degrees) - float64(minutes) /60 ) * 3600)
    return
}
