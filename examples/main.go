package main

import (
    "fmt"
    libGPS "github.com/dantheman213/gps"
)

func main() {
    testCodes := []string{
        "$GPGGA,134658.00,5106.9792,N,11402.3003,W,2,09,1.0,1048.47,M,-16.27,M,08,AAAA*60",
        "$GPGSA,A,3,,,,,,16,18,,22,24,,,3.6,2.1,2.2*3C",
        "$GPRMC,225446.00,A,4916.00,N,12311.12,W,0.5,054.7,060120,,,D*66",
        "$GPGSV,3,1,11,03,03,111,00,04,15,270,00,06,01,010,00,13,06,292,00*74",
        "$GPVTG,360.0,T,348.7,M,0.092,N,0.171,K,D*2A",
    }

    gps := libGPS.NewGPS()
    for _, code := range testCodes {
        gps.IngestNMEASentences(code)
    }

    fmt.Printf("GPS Location: %s\n", gps.GetGPSLocationPretty())
    fmt.Printf("Speed Over Ground: %s km/h\n", gps.NMEA.VTGCourseAndGroundSpeed.SpeedOverGroundKPH)

    fmt.Println("Complete!")
}
