package nmea

import (
	"fmt"
	"math"
)

// PrependXZeroFloat return string with required number of zero (as prefix)
// to have X number before coma
func PrependXZero(value float64, formatString string, expected uint) string {
	rv := fmt.Sprintf(formatString, value)
	nbChar := uint(len(fmt.Sprintf("%d", int(math.Floor(value)))))
	for nbChar < expected {
		rv = "0" + rv
		nbChar += 1
	}
	return rv
}

func PrependToFloatXZero(value float64, expected uint) string {
	return PrependXZero(value, "%f", expected)
}

func PrependToIntXZero(value int, expected uint) string {
	return PrependXZero(float64(value), "%.0f", expected)
}

// Round a float with expected as accuracy
func Round(value float64, expected int) float64 {
	pow := math.Pow(10, float64(expected))
	digit := pow * value
	if _, div := math.Modf(digit); div >= .5 {
		return math.Ceil(digit) / pow
	}
	return math.Floor(digit) / pow
}
