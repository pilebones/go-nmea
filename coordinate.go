package nmea

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const (
	// Allowed cardinal points
	North CardinalPoint = "N"
	South CardinalPoint = "S"
	East  CardinalPoint = "E"
	West  CardinalPoint = "W"
)

var (
	DMFormat = regexp.MustCompile(`^([0-9]*\.?[0-9]+)\ ?([NSEW])?$`)
)

type CardinalPoint string

func (c CardinalPoint) String() string {
	return string(c)
}

func ParseCardinalPoint(raw string) (cp CardinalPoint, err error) {
	cp = CardinalPoint(raw)
	switch cp {
	case North, South, East, West:
	default:
		err = fmt.Errorf("unknow value")
	}
	return
}

const (
	// LatLong Thresholds (ie: spherical degrees)
	// Min is the minimum value allowed for a LatLong
	Min LatLong = -180
	// Max is the maximum value allowed for a LatLong
	Max LatLong = 180
)

type LatLong float64

// NewLatLong parses input has coordinate or return error
//
// Allowed format:
// - DMS (Degrees, Minutes, Secondes), ie: "N 31° 50' 72.38'"
// - DD (Decimal Degree), ie: "31.8534389" "22.870216666666668"
func NewLatLong(raw string) (l LatLong, err error) {

	if raw = strings.TrimSpace(raw); len(raw) == 0 {
		err = fmt.Errorf("Invalid LatLong, can't be empty")
		return
	}

	var value float64
	if value, err = strconv.ParseFloat(raw, 64); err == nil {
		l = LatLong(value)
	} else {
		// Try another format like DM
		if l, err = ParseDM(raw); err != nil {
			return
		}
	}

	if l < Min || l > Max {
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

	matches := DMFormat.FindAllStringSubmatch(strings.TrimSpace(raw), 1)
	if len(matches) == 0 || len(matches[0]) < 3 {
		return LatLong(0), fmt.Errorf("Wrong DM format, got: \"%s\"", raw)
	}

	// Explode data
	if dm, err = strconv.ParseFloat(matches[0][1], 64); err != nil {
		return LatLong(0), err
	}

	if dir, err = ParseCardinalPoint(matches[0][2]); err != nil {
		return LatLong(0), err
	}

	// Compute LatLong
	d := math.Floor(dm / 100) // div dm by 100 and truncate decimal value to get only degrees
	m := dm - (d * 100)       // Sub degrees to dm value
	dm = d + m/60             // switch minute to degree to get value in the same referential

	switch dir {
	case North, East:
		return LatLong(dm), nil
	case South, West:
		return LatLong(0 - dm), nil
	default:
		return 0, fmt.Errorf("Wrong direction (got: %s)", dir.String())
	}
}

// CardinalPoint return the cardinal point related to the kind of coordinate (long or lat)
func (l LatLong) CardinalPoint(isLatitude bool) CardinalPoint {
	if l == 0 {
		return ""
	}

	if l < 0 {
		if isLatitude {
			return South
		}
		return West
	}

	if isLatitude {
		return North
	}
	return East

}

// DM extract degrees and minutes
func (l LatLong) DM() (int, float64) {
	if l < 0 {
		l = 0 - l
	}

	d := math.Floor(float64(l))
	m := (float64(l) - d) * 60

	return int(d), m
}

// DMS extract degrees, minutes and secondes
func (l LatLong) DMS() (int, int, float64) {
	if l < 0 {
		l = 0 - l
	}

	d, m := l.DM()
	m = math.Floor(m)

	// TODO: round secondes
	s := ((float64(l) - (float64(d) + (m / 60))) * 60 * 60)

	return d, int(m), s
}

// Serialize return string like ‘ddmm.mmmm’: degree and minutes as GPS module provide
func (l LatLong) Serialize() string {
	if l == 0 {
		return ""
	}
	d, m := l.DM()
	return strings.Trim(fmt.Sprintf("%d%f", d, m), "0")
}

func (l LatLong) ToDM() string {
	if l == 0 {
		return ""
	}
	return strings.Trim(fmt.Sprintf("%f", l), "0")
}

// PrintDMS return string like: dd° mm' ss.ss" to be human readable
func (l LatLong) ToDMS() string {
	degrees, minutes, secondes := l.DMS()
	return fmt.Sprintf("%d° %d' %.3f\"", degrees, minutes, Round(secondes, 3, .8))
}
