package main

import (
    "fmt"
    libGPS "github.com/dantheman213/gps"
)

// NMEA codes that GPS devices output
var testNMEACodes = [...]string{
    "$GPGGA,061004.114,3404.4083,N,11822.5953,W,1,4,1.9,2.0,M,,,,*2C",
    "$GPGSA,A,3,8,11,15,22,,,,,,,,,3.3,1.9,2.2*06",
    "$GPGLL,061004.114,3404.4083,N,11822.5953,W,1,4,1.9,2.0,M,,,,*2C",
    "$GPRMC,061004.114,A,3404.4083,N,11822.5953,W,10.2,47.0,060120,,,*37",
    "$GPGSV,3,1,11,03,03,111,00,04,15,270,00,06,01,010,00,13,06,292,00*74",
    "$GPVTG,360.0,T,348.7,M,0.092,N,0.171,K,D*2A",
}

func main() {
    gps := libGPS.NewGPS()
    for _, code := range testNMEACodes {
        gps.IngestNMEASentences(code)
    }

    fmt.Printf("GPS Location: %s\n", gps.GetGPSLocationPretty())
    fmt.Printf("Speed Over Ground: %s km/h\n", gps.NMEA.VTGCourseAndGroundSpeed.SpeedOverGroundKPH)

    fmt.Println("Complete!")
}
