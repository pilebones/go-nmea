package nmea

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

type LongLat float64
