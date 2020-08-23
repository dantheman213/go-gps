package gps

import (
    "errors"
    "fmt"
    "github.com/dantheman213/go-gps/location"
    "github.com/dantheman213/go-gps/math"
    "github.com/dantheman213/go-gps/nmea"
    "strconv"
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
    ProviderGPS = "GPSEngine"
    ProviderGLONASS = "GLONASS"
    ProviderGNSS = "GNSS"
)

type GPSEngine struct {
    NMEA *nmea.NMEA
    originalLocation *location.LocationDD
}

func NewGPS() *GPSEngine {
    r := &GPSEngine{}
    r.NMEA = &nmea.NMEA{}
    return r
}

// Get distance traveled in Kilometers
func (g *GPSEngine) GetDistanceTraveledInKM() (float64, error) {
    c1, err1 := g.GetOriginalGPSLocation()
    if err1 != nil {
        return 0, err1
    }

    c2, err2 := g.GetGPSLocation()
    if err2 != nil {
        return 0, err2
    }

    return math.CalculateDistanceBetweenPointsInKM(c1, c2), nil
}

func (g *GPSEngine) GetOriginalGPSLocation() (*location.LocationDD, error) {
    if g.originalLocation == nil {
        return nil, errors.New("location has not been recorded")
    }

    return g.originalLocation, nil
}

func (g *GPSEngine) GetGPSLocation() (*location.LocationDD, error) {
    if g.NMEA.GGALocationFixData != nil {
        f1, err := strconv.ParseFloat(g.NMEA.GGALocationFixData.LatitudeDDM, 64)
        if err != nil {
            return nil, err
        }
        lat, err := math.DDMToDD(f1, g.NMEA.GGALocationFixData.LatitudeDirection)
        if err != nil {
            return nil, err
        }

        f2, err := strconv.ParseFloat(g.NMEA.GGALocationFixData.LongitudeDDM, 64)
        if err != nil {
            return nil, err
        }
        long, err := math.DDMToDD(f2, g.NMEA.GGALocationFixData.LongitudeDirection)
        if err != nil {
            return nil, err
        }

        return &location.LocationDD{
            Latitude:  lat,
            Longitude: long,
        }, nil
    }

    return nil, errors.New("no GGA sentence has been ingested to determine location")
}

func (g *GPSEngine) GetGPSLocationInDDPretty() string {
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

func (g *GPSEngine) GetGPSLocationInDDMPretty() string {
    str := ""
    if g.NMEA.GGALocationFixData != nil {
        str = fmt.Sprintf("%s%s, %s%s", g.NMEA.GGALocationFixData.LatitudeDDM, g.NMEA.GGALocationFixData.LatitudeDirection, g.NMEA.GGALocationFixData.LongitudeDDM, g.NMEA.GGALocationFixData.LongitudeDirection)
    }

    return str
}

func (g *GPSEngine) GetGPSLocationInDMSPretty() string {
    // TODO
    return "TODO"
}

func (g *GPSEngine) GetPrimaryProvider() string {
    list := []int{
        g.NMEA.GPCount,
        g.NMEA.GLCount,
        g.NMEA.GNCount,
    }
    _, max := math.MinMax(list)
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

func (g *GPSEngine) ingestSatelliteNetworkType(prefix string) {
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

func (g *GPSEngine) IngestNMEASentences(sentences string) {
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
                if g.originalLocation == nil {
                    g.originalLocation, err = g.GetGPSLocation()
                    if err != nil {
                        // TODO
                        return
                    }
                }
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
