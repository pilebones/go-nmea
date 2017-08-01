package nmea

import (
	"fmt"
	"testing"
)

func TestNMEAMessage(t *testing.T) {
	msg, err := Parse("$GPGGA,015540.000,3150.68378,N,11711.93139,E,1,17,0.6,0051.6,M,0.0,M,,*58")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(msg)
}
