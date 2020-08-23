package serial

import (
    "errors"
    "fmt"
    libTSerial "github.com/tarm/serial"
    "strings"
    "time"
)

type GPSDevice struct {
    PortName   string
    BaudRate   int
    Port       *libTSerial.Port
}

var buf = make([]byte, 1024) // re-usable / better performance

func Connect(portName string, baudRate int, timeout int) (*GPSDevice, error) {
    c := &libTSerial.Config{
        Name:        portName,
        Baud:        baudRate,
        ReadTimeout: time.Duration(timeout) * time.Second,
        Size:        8,
    }

    p, err := libTSerial.OpenPort(c)
    if err != nil {
        return nil, err
    }

    return &GPSDevice{
        Port:       p,
        PortName:   portName,
        BaudRate:   baudRate,
    }, nil
}

func ReadSerialData(port *libTSerial.Port) (string, error) {
    str := ""
    for true {
        n, err := port.Read(buf)
        if err != nil {
            return "", err
        }
        str += string(buf[:n])

        last := str[len(str)-1:]
        if strings.Index(last, "\r") > -1 || strings.Index(last, "\n") > -1 {
            return fmt.Sprintf("%s\n", strings.TrimSpace(str)), nil
        }
    }

    return "", errors.New("couldn't read serial data")
}
