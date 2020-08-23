package serial

import (
    "errors"
    libDSerial "github.com/dantheman213/go-serial"
    "github.com/dantheman213/go-gps/nmea"
    "log"
)

const (
    dataDetectIterationCount = 10
    portDetectTimeout        = 2
)

var baudRates = [...]int{115200, 38400, 19200, 9600, 4800}

func DetectGPSDevice() (*GPSDevice, error) {
    log.Println("[info] search and auto-detect GPS devices in progress...")

    ports, err := libDSerial.ListPorts()
    if err != nil {
        return nil, err
    }

    for _, currentPort := range *ports {
        for bIndex := 0; bIndex < len(baudRates); bIndex++ {
            // TODO deal with MacOS serial ports with baudrate as suffix
            currentBaudRate := baudRates[bIndex]
            log.Printf("[info] attemping to contact port %s with baud rate %d", currentPort.PortName, currentBaudRate)
            d, err := Connect(currentPort.PortName, baudRates[bIndex], portDetectTimeout)

            if err != nil {
                log.Println("[info] couldn't establish a connection")
                continue
            } else {
                // Found a valid serial device to connect, now check if it's an NMEA compliant GPS serial device
                // Check in multiple iterations to ensure not correct protocol
                for commCount := 0; commCount < dataDetectIterationCount; commCount++ {
                    str, err := ReadSerialData(d.Port)
                    if err != nil {
                        if err.Error() == "EOF" {
                            continue
                        }
                        log.Printf("[info] %s", err)
                    } else {
                        if nmea.IsValidNMEASentence(str) {
                            log.Printf("[info] found GPS device at port: %s baud rate: %d", d.PortName, d.BaudRate)
                            return d, nil
                        }
                    }
                }

                log.Println("[info] found data but doesn't appear to be NMEA... skipping...")
            }

            log.Printf("[info] closing port %s", currentPort.PortName)
            _ = d.Port.Close()
        }
    }

    return nil, errors.New("could not find a valid GPS serial device")
}
