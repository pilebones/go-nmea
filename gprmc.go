package nmea

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Examples:
// $GPRMC,013732.000,A,3150.7238,N,11711.7278,E,0.00,0.00,220413,,,A*68
// $GPRMC,081836,A,3751.65,S,14507.36,E,000.0,360.0,130998,011.3,E*62
// $GPRMC,225446,A,4916.45,N,12311.12,W,000.5,054.7,191194,020.3,E*68
// $GPRMC,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*70

func NewGPRMC(m Message) *GPRMC {
	return &GPRMC{Message: m}
}

type GPRMC struct {
	Message

	DateTimeUTC       time.Time // Aggregation of TimeUTC+Date data field
	IsValid           DataValid // 'V' =Invalid / 'A' = Valid
	Latitude          LatLong   // In decimal format
	Longitude         LatLong   // In decimal format
	Speed             float64   // Speed over ground in knots
	COG               float64   // Course over ground in degree
	MagneticVariation float64   // Magnetic variation in degree, not being output
	PositioningMode   PositioningMode
}

func (m *GPRMC) GetMessage() *Message { // Implement NMEA interface
	return &m.Message
}

func (m *GPRMC) parse() (err error) {
	if len(m.Fields) != 12 {
		return fmt.Errorf("Incomplete GPRMC message, not enougth data fields (got: %d, wanted: %d)", len(m.Fields), 12)
	}

	datetime := fmt.Sprintf("%s %s", m.Fields[8], m.Fields[0])
	if m.DateTimeUTC, err = time.Parse("020106 150405.000", datetime); err != nil {
		return fmt.Errorf("Unable to parse datetime UTC from data field (got: %s)", datetime)
	}

	m.IsValid = (m.Fields[1] == "A")

	if m.Latitude, err = NewLatLong(strings.Join(m.Fields[2:4], " ")); err != nil {
		return err
	}
	if m.Longitude, err = NewLatLong(strings.Join(m.Fields[4:6], " ")); err != nil {
		return err
	}

	if m.Speed, err = strconv.ParseFloat(m.Fields[6], 64); err != nil {
		return fmt.Errorf("Unable to parse speed from data field (got: %s)", m.Fields[6])
	}

	if m.COG, err = strconv.ParseFloat(m.Fields[7], 64); err != nil {
		return fmt.Errorf("Unable to parse course over ground from data field (got: %s)", m.Fields[7])
	}

	if len(m.Fields[9]) > 0 {
		if m.MagneticVariation, err = strconv.ParseFloat(m.Fields[9], 64); err != nil {
			return fmt.Errorf("Unable to parse magnetic variation from data field (got: %s)", m.Fields[9])
		}

		if len(m.Fields[10]) > 0 {
			magneticVariationDir, err := ParseCardinalPoint(m.Fields[10])
			if err != nil {
				return fmt.Errorf("Unable to parse magnetic variation indicator from data field (got: %s)", m.Fields[10])
			}

			switch magneticVariationDir {
			case WEST:
				m.MagneticVariation = 0 - m.MagneticVariation
			case EAST:
				// Allowed direction
			default:
				return fmt.Errorf("Wrong magnetic variation direction (got: %s)", m.Fields[10])
			}
		}
	}

	if m.PositioningMode, err = ParsePositioningMode(m.Fields[11]); err != nil {
		return fmt.Errorf("Unable to parse GPS positioning mode from data field (got: %s)", m.Fields[11])
	}

	return nil
}
