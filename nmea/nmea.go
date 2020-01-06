package nmea

import (
    "errors"
    "strings"
)

type NMEA struct {
    GPCount                  int // GPS
    GLCount                  int // GLONASS
    GNCount                  int // GNSS
    GGALocationFixData       *GGA
    RMCRecMinData            *RMC
    GSAOverallSatelliteData  *GSA
    GSVDetailedSatelliteData *[]GSV
    VTGCourseAndGroundSpeed  *VTG
}

type GGA struct {
    Timestamp          string
    LatitudeDDM        string
    LatitudeDirection  string
    LongitudeDDM       string
    LongitudeDirection string
    FixQuality         string
    Satellites         string
    Checksum           string
}

type RMC struct {
    Timestamp              string
    SpeedOverGroundInKnots string
    TrackAngleInDegrees    string
    Date                   string
    MagneticVariation      string
    Checksum               string
}

type GSA struct {
    Mode1    string
    Mode2    string
    PDOP     string
    HDOP     string
    VDOP     string
    Checksum string
}

// SV = Satellite Vehicle
type GSV struct {
    VisibleSVCount              string
    MessageNumber               string // 1-3
    MessageCountInCycle         string // 1-3
    SVPRN                       string
    ElevationDegrees            string
    AzimuthDegreesFromTrueNorth string
    SNR                         string
    Checksum                    string
}

type VTG struct {
    TrackMadeGoodDegreesTrue     string
    TrackMadeGoodDegreesMagnetic string
    SpeedInKnots                 string
    SpeedOverGroundKPH           string
    Checksum                     string
}

func (g *GGA) GetLatitudeDDM() (string, string, error) {
    return g.LatitudeDDM, g.LatitudeDirection, nil
}

func (g *GGA) GetLongitudeDDM() (string, string, error) {
    return g.LongitudeDDM, g.LongitudeDirection, nil
}

func ParseGGA(s string) (*GGA, error) {
    tokens := strings.Split(s, ",")
    if len(tokens) >= 15 {
        return &GGA{
            Timestamp:          tokens[1],
            LatitudeDDM:        tokens[2],
            LatitudeDirection:  tokens[3],
            LongitudeDDM:       tokens[4],
            LongitudeDirection: tokens[5],
            FixQuality:         tokens[6],
            Satellites:         tokens[7],
            Checksum:           tokens[14][strings.Index(tokens[14], "*")+1 : len(tokens[14])],
        }, nil
    }

    return nil, errors.New("malformed NMEA sentence")
}

func ParseRMC(s string) (*RMC, error) {
    tokens := strings.Split(s, ",")
    if len(tokens) >= 12 {
        return &RMC{
            Timestamp:              tokens[1],
            SpeedOverGroundInKnots: tokens[7],
            TrackAngleInDegrees:    tokens[8],
            Date:                   tokens[9],
            MagneticVariation:      tokens[10],
            Checksum:               tokens[12][strings.Index(tokens[12], "*")+1 : len(tokens[12])],
        }, nil
    }

    return nil, errors.New("malformed NMEA sentence")
}

func ParseGSA(s string) (*GSA, error) {
    tokens := strings.Split(s, ",")
    i := strings.Index(tokens[17], "*")
    if len(tokens) >= 18 {
        return &GSA{
            Mode1:    tokens[1],
            Mode2:    tokens[2],
            PDOP:     tokens[15],
            HDOP:     tokens[16],
            VDOP:     tokens[17][0:i],
            Checksum: tokens[17][i+1 : len(tokens[17])],
        }, nil
    }

    return nil, errors.New("malformed NMEA sentence")
}

func ParseGSV(s string) (*GSV, error) {
    //tokens := strings.Split(s, ",")
    //if len(tokens) >= 8 {
    //    return &GSV{
    //        VisibleSVCount:              "",
    //        MessageNumber:               "",
    //        MessageCountInCycle:         "",
    //        SVPRN:                       "",
    //        ElevationDegrees:            "",
    //        AzimuthDegreesFromTrueNorth: "",
    //        SNR:                         "",
    //        Checksum:                    "",
    //    }
    //}
    //
    //return nil, errors.New("malformed NMEA sentence")
    return nil, nil
}

func ParseVTG(s string) (*VTG, error) {
    tokens := strings.Split(s, ",")
    if len(tokens) >= 10 {
        return &VTG{
            TrackMadeGoodDegreesTrue:     tokens[1],
            TrackMadeGoodDegreesMagnetic: tokens[3],
            SpeedInKnots:                 tokens[5],
            SpeedOverGroundKPH:           tokens[7],
            Checksum:                     tokens[9][strings.Index(tokens[9], "*")+1 : len(tokens[9])],
        }, nil
    }

    return nil, errors.New("malformed NMEA sentence")
}

func IsValidNMEASentence(sentence string) bool {
    if len(sentence) > 7 {
        if sentence[0:2] == "$G" {
            firstCommaPos := strings.Index(sentence, ",")
            if firstCommaPos == 6 || firstCommaPos == 7 || firstCommaPos == 9 {
                // ex: $GPGGA, $GPPTNL, $GPPFUGDP
                return true
            }
        }
    }

    return false
}
