package nmea

import "fmt"

const (
	NORTH CardinalPoint = "N"
	SOUTH CardinalPoint = "S"
	EAST  CardinalPoint = "E"
	WEST  CardinalPoint = "W"
)

type CardinalPoint string

func (c CardinalPoint) String() string {
	return string(c)
}

func ParseCardinalPoint(raw string) (*CardinalPoint, error) {
	cp := CardinalPoint(raw)
	switch cp {
	case NORTH, SOUTH, EAST, WEST:
		return &cp, nil
	default:
		return nil, fmt.Errorf("unknow value")
	}
}

const (
	MIN_LATLONG_THRESHOLD LatLong = -180
	MAX_LATLONG_THRESHOLD LatLong = 180
)

type LatLong float64

// NewLatLong parses input has coordinate or return error
//
// Allowed format:
// - DMS (Degrees, Minutes, Secondes), ie: "N 31° 50' 72.38'"
// - DD (Decimal Degree), ie: "31.8534389" "22.870216666666668"
func NewLatLong(raw string) (l LatLong, err error) {
	if l, err = ParseDMS(raw); err == nil {
		return
	}

	if l, err = ParseDD(raw); err == nil {
		return
	}

	if l < MIN_LATLONG_THRESHOLD || l > MAX_LATLONG_THRESHOLD {
		return l, fmt.Errorf("invalid range (got: %f)", l)
	}

	return LatLong(0), fmt.Errorf("invalid format (got: %s)", raw)
}

func ParseDD(raw string) (l LatLong, err error) {
	return
}

func ParseDMS(raw string) (l LatLong, err error) {
	return
}
