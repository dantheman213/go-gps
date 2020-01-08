package gps

import (
    "errors"
    "fmt"
    "github.com/dantheman213/gps/convert"
    "github.com/dantheman213/gps/internal/common"
    "github.com/dantheman213/gps/nmea"
    "strings"
)

const (
    DirectionNorth     = "N"
    DirectionEast      = "E"
    DirectionSouth     = "S"
    DirectionWest      = "W"
    DirectionNorthEast = "NE"
    DirectionNorthWest = "NW"
    DirectionSouthEast = "SE"
    DirectionSouthWest = "SW"
    ProviderGPS = "GPS"
    ProviderGLONASS = "GLONASS"
    ProviderGNSS = "GNSS"
)

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

type GPS struct {
    NMEA *nmea.NMEA
}

func NewGPS() *GPS {
    r := &GPS{}
    r.NMEA = &nmea.NMEA{}
    return r
}

func (g *GPS) GetGPSLocation() (*LocationDD, error) {
    if g.NMEA.GGALocationFixData != nil {
        lat, err := convert.DDMToDD(g.NMEA.GGALocationFixData.LatitudeDDM, g.NMEA.GGALocationFixData.LatitudeDirection)
        if err != nil {
            return nil, err
        }

        long, err := convert.DDMToDD(g.NMEA.GGALocationFixData.LongitudeDDM, g.NMEA.GGALocationFixData.LongitudeDirection)
        if err != nil {
            return nil, err
        }

        return &LocationDD{
            Latitude:  lat,
            Longitude: long,
        }, nil
    }

    return nil, errors.New("no GGA sentence has been ingested to determine location")
}

func (g *GPS) GetGPSLocationInDDPretty() string {
    loc, err := g.GetGPSLocation()
    if err != nil {
        // TODO
    }

    str := ""
    if loc != nil {
        str = fmt.Sprintf("%f, %f", loc.Latitude, loc.Longitude)
    }

    return str
}

func (g *GPS) GetGPSLocationInDDMPretty() string {
    str := ""
    if g.NMEA.GGALocationFixData != nil {
        str = fmt.Sprintf("%s%s, %s%s", g.NMEA.GGALocationFixData.LatitudeDDM, g.NMEA.GGALocationFixData.LatitudeDirection, g.NMEA.GGALocationFixData.LongitudeDDM, g.NMEA.GGALocationFixData.LongitudeDirection)
    }

    return str
}

func (g *GPS) GetGPSLocationInDMSPretty() string {
    // TODO
    return "TODO"
}

func (g *GPS) GetPrimaryProvider() string {
    list := []int{
        g.NMEA.GPCount,
        g.NMEA.GLCount,
        g.NMEA.GNCount,
    }
    _, max := common.MinMax(list)
    switch max {
    case g.NMEA.GPCount:
        return ProviderGPS
    case g.NMEA.GLCount:
        return ProviderGLONASS
    case g.NMEA.GNCount:
        return ProviderGNSS
    }

    return ""
}

func (g *GPS) ingestSatelliteNetworkType(prefix string) {
    switch prefix {
    case "GP":
        g.NMEA.GPCount = g.NMEA.GPCount + 1
        break
    case "GL":
        g.NMEA.GLCount += 1
        break
    case "GN":
        g.NMEA.GNCount += 1
        break
    default:
        // TODO
    }
}

func (g *GPS) IngestNMEASentences(sentences string) {
    s := strings.ReplaceAll(sentences, "\r", "")
    items := strings.Split(s, "\n")

    for _, item := range items {
        if nmea.IsValidNMEASentence(item) {
            g.ingestSatelliteNetworkType(item[1:3])

            nmeaCode := item[3:strings.Index(item, ",")]
            switch nmeaCode {
            case "GGA":
                d, err := nmea.ParseGGA(item)
                if err != nil {
                    // TODO
                    return
                }
                g.NMEA.GGALocationFixData = d
                break
            case "RMC":
                d, err := nmea.ParseRMC(item)
                if err != nil {
                    // TODO
                    return
                }
                g.NMEA.RMCRecMinData = d
                break
            case "GSA":
                d, err := nmea.ParseGSA(item)
                if err != nil {
                    // TODO
                    return
                }
                g.NMEA.GSAOverallSatelliteData = d
                break
            case "GSV":
                // TODO
                break
            case "VTG":
                d, err := nmea.ParseVTG(item)
                if err != nil {
                    // TODO
                    return
                }
                g.NMEA.VTGCourseAndGroundSpeed = d
                break
            default:
                // TODO ?
            }
        } else {
            // unsupported or invalid NMEA sentence
            // TODO
        }
    }
}
