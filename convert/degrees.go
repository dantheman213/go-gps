package convert

import (
    "errors"
    "math"
    "strconv"
)

// Parse DDM (Degrees Decimal Minutes) to DD (Decimal Degrees)
func ConvertDDMToDD(value string, direction string) (float32, error) {
    if value == "" || direction == "" {
        return 0, errors.New("the location and / or direction value does not exist")
    }

    lat, _ := strconv.ParseFloat(value, 64)
    degrees := math.Floor(lat / 100)
    minutes := ((lat / 100) - math.Floor(lat/100)) * 100 / 60
    decimal := degrees + minutes

    if direction == "W" || direction == "S" {
        decimal *= -1
    }

    return float32(decimal), nil
}
