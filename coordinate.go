package nmea

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

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

func ParseCardinalPoint(raw string) (cp CardinalPoint, err error) {
	cp = CardinalPoint(raw)
	switch cp {
	case NORTH, SOUTH, EAST, WEST:
	default:
		err = fmt.Errorf("unknow value")
	}
	return
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
	if l, err = ParseDM(raw); err != nil {
		return
	}

	if l < MIN_LATLONG_THRESHOLD || l > MAX_LATLONG_THRESHOLD {
		err = fmt.Errorf("invalid range (got: %f)", l)
	}
	return
}

// ParseDM return LatLong from provided format from GPS module (in format ‘ddmm.mmmm’: degree and minutes)
// Allowed format: "3150.7238N" or "3150.7238 N"
// @see https://fr.wikipedia.org/wiki/Coordonn%C3%A9es_g%C3%A9ographiques
// => 1 degree = 60 minutes
// => 1 minute = 60 secondes
// Example: Baltimore (United state) => latitude = 39,28° N, longitude = 76,60° O (39° 17′ N, 76° 36′ O).
// 0.28° = (0.28°*60min)/1° = 16.8min => ~17 minutes
func ParseDM(raw string) (LatLong, error) {

	var (
		dir CardinalPoint
		dm  float64
		err error
	)

	// Explode data
	if dm, err = strconv.ParseFloat(strings.TrimSpace(string(raw[:len(raw)-2])), 64); err != nil {
		return LatLong(0), err
	}

	if dir, err = ParseCardinalPoint(string(raw[len(raw)-1])); err != nil {
		return LatLong(0), err
	}

	// Compute LatLong
	d := math.Floor(dm / 100) // div dm by 100 and truncate decimal value to get only degrees
	m := dm - (d * 100)       // Sub degrees to dm value
	dm = d + m/60             // switch minute to degree to get value in the same referential

	switch dir {
	case NORTH, EAST:
		return LatLong(dm), nil
	case SOUTH, WEST:
		return LatLong(0 - dm), nil
	default:
		return 0, fmt.Errorf("Wrong direction (got: %s)", dir.String())
	}
}

func (l LatLong) PrintDMS() string {
	degrees := math.Floor(float64(l))
	minutes := math.Floor((float64(l) - degrees) * 60)
	secondes := (float64(l) - (degrees + (minutes / 60))) * 60 * 60 // TODO: round secondes
	return fmt.Sprintf("%d° %d' %f\"", int(degrees), int(minutes), secondes)
}
