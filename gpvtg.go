package nmea

import (
	"fmt"
	"strconv"
)

// Examples:
// $GPVTG,0.0,T,,M,0.0,N,0.1,K,A*0C

func NewGPVTG(m Message) *GPVTG {
	return &GPVTG{Message: m}
}

type GPVTG struct {
	Message

	COG             float64 // Course over ground (true) in degree
	SpeedKnots      float64 // Speed over ground in knots
	SpeedKmh        float64 // Speed over ground in km/h
	PositioningMode PositioningMode
}

func (m *GPVTG) parse() (err error) {
	if len(m.Fields) != 9 {
		return m.Error(fmt.Errorf("Incomplete GPVTG message, not enougth data fields (got: %d, wanted: %d)", len(m.Fields), 9))
	}

	// Validate fixed field
	for i, v := range map[int]string{1: "T", 3: "M", 5: "N", 7: "K"} {
		if m.Fields[i] != v {
			return m.Error(fmt.Errorf("Invalid fixed field at %d (got: %s, wanted: %s)", i+1, m.Fields[i], v))
		}
	}

	if m.COG, err = strconv.ParseFloat(m.Fields[0], 64); err != nil {
		return m.Error(fmt.Errorf("Unable to parse true course over ground from data field (got: %s)", m.Fields[0]))
	}

	if m.SpeedKnots, err = strconv.ParseFloat(m.Fields[4], 64); err != nil {
		return m.Error(fmt.Errorf("Unable to parse speed from data field (got: %s)", m.Fields[4]))
	}

	if m.SpeedKmh, err = strconv.ParseFloat(m.Fields[6], 64); err != nil {
		return m.Error(fmt.Errorf("Unable to parse speed from data field (got: %s)", m.Fields[6]))
	}

	if m.PositioningMode, err = ParsePositioningMode(m.Fields[8]); err != nil {
		return m.Error(fmt.Errorf("Unable to parse GPS positioning mode from data field (got: %s)", m.Fields[8]))
	}

	return nil
}
