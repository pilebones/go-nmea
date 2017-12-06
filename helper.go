package nmea

import (
	"fmt"
	"math"
)

// PrependXZero return string with expected number of zero (as prefix)
// to have X number before coma
func PrependXZero(value float64, formatString string, expected uint) string {
	rv := fmt.Sprintf(formatString, value)
	nbChar := uint(len(fmt.Sprintf("%d", int(math.Floor(value)))))
	for nbChar < expected {
		rv = "0" + rv
		nbChar++
	}
	return rv
}

// PrependToFloatXZero doing same thing that PrependXZeroFloat but with float as input
func PrependToFloatXZero(value float64, expected uint) string {
	return PrependXZero(value, "%f", expected)
}

// PrependToIntXZero doing same thing that PrependXZeroFloat but with int as input
func PrependToIntXZero(value int, expected uint) string {
	return PrependXZero(float64(value), "%.0f", expected)
}

// Round a float with expected as accuracy
func Round(value float64, expected int, threshold float64) float64 {
	pow := math.Pow(10, float64(expected))
	digit := pow * value
	if _, div := math.Modf(digit); div >= threshold {
		return math.Ceil(digit) / pow
	}
	return math.Floor(digit) / pow
}
