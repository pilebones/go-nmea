package nmea

import (
	"fmt"
	"strings"
	"time"
)

// Examples:
// $GPGLL,3110.2908,N,12123.2348,E,041139.000,A,A*59

func NewGPGLL(m Message) *GPGLL {
	return &GPGLL{Message: m}
}

type GPGLL struct {
	Message

	TimeUTC             time.Time // Aggregation of TimeUTC data field
	Latitude, Longitude LatLong   // In decimal format
	IsValid             DataValid
	PositioningMode     PositioningMode
}

func (m *GPGLL) GetMessage() *Message { // Implement NMEA interface
	return &m.Message
}

func (m *GPGLL) parse() (err error) {
	if len(m.Fields) != 7 {
		return m.Error(fmt.Errorf("Incomplete GPGLL message, not enougth data fields (got: %d, wanted: %d)", len(m.Fields), 7))
	}

	if m.Latitude, err = NewLatLong(strings.Join(m.Fields[0:2], " ")); err != nil {
		return m.Error(err)
	}
	if m.Longitude, err = NewLatLong(strings.Join(m.Fields[2:4], " ")); err != nil {
		return m.Error(err)
	}

	if m.TimeUTC, err = time.Parse("150405.000", m.Fields[4]); err != nil {
		return m.Error(fmt.Errorf("Unable to parse time UTC from data field (got: %s)", m.Fields[4]))
	}

	m.IsValid = (m.Fields[5] == "A")

	if m.PositioningMode, err = ParsePositioningMode(m.Fields[6]); err != nil {
		return m.Error(fmt.Errorf("Unable to parse GPS positioning mode from data field (got: %s)", m.Fields[6]))
	}

	return nil
}
