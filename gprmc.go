package nmea

import "time"

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

	TimeUTC                    time.Time
	Valid                      RMCValid // 'V' =Invalid / 'A' = Valid
	Latitude                   LongLat
	Longitude                  LongLat
	Speed                      float64       // Speed over ground in knots
	COG                        float64       // Course over ground in degree
	Date                       time.Time     // Date
	MagneticVariation          float64       // Magnetic variation in degree, not being output
	MagneticVariationIndicator CardinalPoint // E/W Magnetic variation E/W indicator, not being output
	PositionningMode           PositionningMode
}

func (s *GPRMC) parse() error {

	return nil
}

const (
	Valid   RMCValid = true
	Invalid RMCValid = false
)

type RMCValid bool

func (v RMCValid) String() string {
	if v == Valid {
		return "A"
	}
	return "V"
}

const (
	NO_FIX                PositionningMode = "N"
	AUTONOMOUS_GNSS_FIX   PositionningMode = "A"
	DIFFERENTIAL_GNSS_FIX PositionningMode = "D"
)

type PositionningMode string
