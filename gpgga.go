package nmea

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Examples:
// $GPGGA,015540.000,3150.68378,N,11711.93139,E,1,17,0.6,0051.6,M,0.0,M,,*58

func NewGPGGA(m Message) *GPGGA {
	return &GPGGA{Message: m}
}

type GPGGA struct {
	Message

	TimeUTC            time.Time // Aggregation of TimeUTC data field
	Latitude           LatLong   // In decimal format
	Longitude          LatLong   // In decimal format
	QualityIndicator   QualityIndicator
	NbOfSatellitesUsed uint64
	HDOP               float64
	Altitude           float64
	GeoIdSep           float64

	// FIXME: Manage field below when I found a sample with no-empty data
	// DGPSAge        *uint64
	// DGPSiStationId *string
}

func (m *GPGGA) GetMessage() *Message { // Implement NMEA interface
	return &m.Message
}

func (m *GPGGA) parse() (err error) {
	if len(m.Fields) != 14 {
		return fmt.Errorf("Incomplete GPGGA message, not enougth data fields (got: %d, wanted: %d)", len(m.Fields), 14)
	}

	// Validate fixed field
	for i, v := range map[int]string{9: "M", 11: "M"} {
		if m.Fields[i] != v {
			return fmt.Errorf("Invalid fixed field at %d (got: %s, wanted: %s)", i+1, m.Fields[i], v)
		}
	}

	if m.TimeUTC, err = time.Parse("150405.000", m.Fields[0]); err != nil {
		return m.Error(fmt.Errorf("Unable to parse time UTC from data field (got: %s)", m.Fields[0]))
	}

	if m.Latitude, err = NewLatLong(strings.Join(m.Fields[1:3], " ")); err != nil {
		return m.Error(err)
	}

	if m.Longitude, err = NewLatLong(strings.Join(m.Fields[3:5], " ")); err != nil {
		return m.Error(err)
	}

	if m.QualityIndicator, err = ParseQualityIndicator(m.Fields[5]); err != nil {
		return m.Error(err)
	}

	if m.NbOfSatellitesUsed, err = strconv.ParseUint(m.Fields[6], 10, 0); err != nil {
		return m.Error(err)
	}

	if m.HDOP, err = strconv.ParseFloat(m.Fields[7], 64); err != nil {
		return m.Error(err)
	}
	if m.Altitude, err = strconv.ParseFloat(m.Fields[8], 64); err != nil {
		return m.Error(err)
	}

	if m.GeoIdSep, err = strconv.ParseFloat(m.Fields[10], 64); err != nil {
		return m.Error(err)
	}

	return nil
}

const (
	QUALITY_INDICATOR_INVALID = iota
	QUALITY_INDICATOR_GNSSS
	QUALITY_INDICATOR_DGPS
)

type QualityIndicator int

func (s QualityIndicator) String() string {
	switch s {
	case QUALITY_INDICATOR_INVALID:
		return "invalid"
	case QUALITY_INDICATOR_GNSSS:
		return "GNSS fix"
	case QUALITY_INDICATOR_DGPS:
		return "DGPS fix"
	default:
		return "unknow"

	}
}

func ParseQualityIndicator(raw string) (qi QualityIndicator, err error) {
	i, err := strconv.ParseInt(raw, 10, 0)
	if err != nil {
		return
	}

	qi = QualityIndicator(i)
	switch qi {
	case QUALITY_INDICATOR_INVALID, QUALITY_INDICATOR_GNSSS, QUALITY_INDICATOR_DGPS:
	default:
		err = fmt.Errorf("unknow value")
	}
	return
}
