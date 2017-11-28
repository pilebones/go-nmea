package nmea

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type latlong struct {
	Latitude  string
	Longitude string
}

type sample struct {
	Name string
	DD   latlong
	DMS  latlong
}

func TestGPSCoordinate(t *testing.T) {
	// Coordinate: Latitude, Longitude

	samples := []sample{
		sample{
			Name: "New York, United States",
			DD: latlong{
				Latitude:  "40.7127753",
				Longitude: "-74.0059728",
			},
			DMS: latlong{
				Latitude:  "N 40° 42' 45.991\"",
				Longitude: "W 74° 0' 21.502\"",
			},
		},
		sample{
			Name: "Lyon, France",
			DD: latlong{
				Latitude:  "45.764043",
				Longitude: "4.835658999999964",
			},
			DMS: latlong{
				Latitude:  "N 45° 45' 50.555\"",
				Longitude: "E 4° 50' 8.372\"",
			},
		},
		sample{
			Name: "Buenos Aires, Argentina",
			DD: latlong{
				Latitude:  "-34.60368440000001",
				Longitude: "-58.381559100000004",
			},
			DMS: latlong{
				Latitude:  "S 34° 36' 13.264\"",
				Longitude: "W 58° 22' 53.612\"",
			},
		},
		sample{
			Name: "Auckland, New Zealand",
			DD: latlong{
				Latitude:  "-36.8484597",
				Longitude: "174.76333150000005",
			},
			DMS: latlong{
				Latitude:  "S 36° 50' 54.455\"",
				Longitude: "E 174° 45' 47.993\"",
			},
		},
		sample{
			Name: "Equatorial Guinea",
			DD: latlong{
				Latitude:  "0.5800767981271677",
				Longitude: "9.755859375",
			},
			DMS: latlong{
				Latitude:  "N 0° 34' 48.276\"",
				Longitude: "E 9° 45' 21.093\"",
			},
		},
	}

	for _, s := range samples {
		lat, err := NewLatLong(s.DD.Latitude)
		if err != nil {
			t.Fatalf("[%s] Invalid lat, err: %v", s.Name, err)
		}
		latDM, _ := strconv.ParseFloat(s.DD.Latitude, 64)
		expectedLatDM := strings.Trim(fmt.Sprintf("%f", latDM), "0")
		if lat.ToDM() != expectedLatDM {
			t.Fatalf("[%s] Wrong latitude conversion to DM format, (got: %s, expected: %s)", s.Name, lat.ToDM(), expectedLatDM)
		}

		/*if lat.ToDMS() != s.DMS.Latitude {
			t.Fatalf("[%s] Wrong conversion to DMS format, (got: %s, expected: %s)", s.Name, lat.ToDMS(), s.DMS.Latitude)
		}*/

		long, err := NewLatLong(s.DD.Longitude)
		if err != nil {
			t.Fatalf("[%s] Invalid long, err: %v", s.Name, err)
		}

		longDM, _ := strconv.ParseFloat(s.DD.Longitude, 64)
		expectedLongDM := strings.Trim(fmt.Sprintf("%f", longDM), "0")
		if long.ToDM() != expectedLongDM {
			t.Fatalf("[%s] Wrong longitude conversion to DM format, (got: %s, expected: %s)", s.Name, long.ToDM(), expectedLongDM)
		}

		/*if long.ToDMS() != s.DMS.Longitude {
			t.Fatalf("[%s] Wrong conversion to DMS format, (got: %s, expected: %s)", s.Name, long.ToDMS(), s.DMS.Latitude)
		}*/
	}
}
