package nmea

import "testing"

func TestNMEAMessage(t *testing.T) {

	nmeas := []string{
		// Common NMEA Packet Protocol
		"$GPGGA,015540.000,3150.68378,N,11711.93139,E,1,17,0.6,0051.6,M,0.0,M,,*58",
		"$GPRMC,013732.000,A,3150.7238,N,11711.7278,E,0.00,0.00,220413,,,A*68",
		"$GPVTG,0.0,T,,M,0.0,N,0.1,K,A*0C",
		"$GPGSA,A,3,14,06,16,31,23,,,,,,,,1.66,1.42,0.84*0F",
		"$GPGSV,3,1,12,01,05,060,18,02,17,259,43,04,56,287,28,09,08,277,28*77",
		"$GPGSV,3,2,12,10,34,195,46,13,08,125,45,17,67,014,,20,32,048,24*74",
		"$GPGSV,3,3,12,23,13,094,48,24,04,292,24,28,49,178,46,32,06,037,22*7D",
		"$GPGLL,3110.2908,N,12123.2348,E,041139.000,A,A*59",
		"$GPTXT,01,01,02,ANTSTATUS=OK*3B",

		// MTK NMEA Packet Protocol
		// From "L80 GPS Protocol Specification"

		// TODO: Manage MTK Packet
		/*
			"$PMTK010,001*2E",
			"$PMTK011,MTKGPS*08",
			"$PMTK001, 869,3*37",
			"$PMTK101*32",
			"$PMTK102*31",
			"$PMTK103*30",
			"$PMTK104*37",
			"$PMTK161,0*28",
			"$PMTK183*38",
			"$PMTKLOG,456,0,11,31,2,0,0,0,3769,46*48",
			"$PMTK184,1*22",
			"$PMTK185,1*23",
			"$PMTK622,1*29",
			"$PMTK225,8*23",
			"$PMTK251,38400*27",
			"$PMTK286,0*22",
			"$PMTK300,1000,0,0,0,0*1C",
			"$PMTK301,2*2E",
			"$PMTK313,1*2E",
		*/
	}

	for _, n := range nmeas {
		if _, err := Parse(n); err != nil {
			t.Fatal(err)
		}
	}
}
