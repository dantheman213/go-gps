package main

import (
    "fmt"
    libGPS "github.com/dantheman213/go-gps"
    "github.com/dantheman213/go-gps/serial"
    "log"
    "strings"
)

func main() {
    device, err := serial.DetectGPSDevice()
    if err != nil {
        log.Fatalf("[error] %s", err)
    }

    gps := libGPS.NewGPS()
    for true {
        dat, err := serial.ReadSerialData(device.Port)
        if err != nil {
            log.Printf("couldn't read data stream")
            return
        }

        gps.IngestNMEASentences(dat)

        fmt.Print(strings.TrimSpace(dat))
        fmt.Println(gps.GetGPSLocationInDDPretty())
    }
}
