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
	GeoIDSep           *float64

	// FIXME: Manage field below when I found a sample with no-empty data
	// DGPSAge        *uint64
	// DGPSiStationId *string
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

	if latitude := strings.TrimSpace(strings.Join(m.Fields[1:3], " ")); len(latitude) > 0 {
		if m.Latitude, err = NewLatLong(latitude); err != nil {
			return m.Error(err)
		}
	}

	if longitude := strings.TrimSpace(strings.Join(m.Fields[3:5], " ")); len(longitude) > 0 {
		if m.Longitude, err = NewLatLong(longitude); err != nil {
			return m.Error(err)
		}
	}

	if m.QualityIndicator, err = ParseQualityIndicator(m.Fields[5]); err != nil {
		return m.Error(err)
	}

	if m.NbOfSatellitesUsed, err = strconv.ParseUint(m.Fields[6], 10, 0); err != nil {
		return m.Error(err)
	}

	if hdop := m.Fields[7]; len(hdop) > 0 {
		if m.HDOP, err = strconv.ParseFloat(hdop, 64); err != nil {
			return m.Error(err)
		}
	}
	if altitude := m.Fields[8]; len(altitude) > 0 {
		if m.Altitude, err = strconv.ParseFloat(altitude, 64); err != nil {
			return m.Error(err)
		}
	}

	if geoIDSep := m.Fields[10]; len(geoIDSep) > 0 {
		id, err := strconv.ParseFloat(geoIDSep, 64)
		if err != nil {
			return m.Error(err)
		}
		m.GeoIDSep = &id
	}

	return nil
}

func (m GPGGA) Serialize() string { // Implement NMEA interface

	hdr := TypeIDs["GPGGA"]
	fields := make([]string, 0)

	fields = append(fields, m.TimeUTC.Format("150405.000"),
		strings.Trim(m.Latitude.Serialize(), "0"), m.Latitude.CardinalPoint(true).String(),
		strings.Trim(m.Longitude.Serialize(), "0"), m.Longitude.CardinalPoint(false).String(),
		strconv.Itoa(int(m.QualityIndicator)),
		strconv.Itoa(int(m.NbOfSatellitesUsed)),
	)

	if m.HDOP > 0 {
		fields = append(fields, fmt.Sprintf("%.1f", m.HDOP))
	} else {
		fields = append(fields, "")
	}

	if m.Altitude > 0 {
		fields = append(fields, PrependXZero(m.Altitude, "%.1f", 4))

	} else {
		fields = append(fields, "")
	}

	fields = append(fields, "M")

	if m.GeoIDSep != nil {
		fields = append(fields, fmt.Sprintf("%.1f", *m.GeoIDSep))
	} else {
		fields = append(fields, "")
	}

	fields = append(fields,
		"M",
		"", // DGPSAge always empty ?
		"", // DGPSiStationId always empty ?
	)

	msg := Message{Type: hdr, Fields: fields}
	msg.Checksum = msg.ComputeChecksum()

	return msg.Serialize()
}

const (
	InvalidIndicator = iota
	GNSSS
	DGPS
)

type QualityIndicator int

func (s QualityIndicator) String() string {
	switch s {
	case InvalidIndicator:
		return "invalid"
	case GNSSS:
		return "GNSS fix"
	case DGPS:
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
	case InvalidIndicator, GNSSS, DGPS:
	default:
		err = fmt.Errorf("unknow value")
	}
	return
}
