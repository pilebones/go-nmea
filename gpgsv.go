package nmea

import "fmt"

// Examples:
// $GPGSV,3,1,12,01,05,060,18,02,17,259,43,04,56,287,28,09,08,277,28*77
// $GPGSV,3,2,12,10,34,195,46,13,08,125,45,17,67,014,,20,32,048,24*74
// $GPGSV,3,3,12,23,13,094,48,24,04,292,24,28,49,178,46,32,06,037,22*7D

func NewGPGSV(m Message) *GPGSV {
	return &GPGSV{Message: m}
}

type Satellite struct {
	Id        string
	Elevation int  // Elevation in degree (0 ~ 90)
	Azimuth   int  // Azimuth in degree (0 ~ 359)
	SNR       *int // Signal to Noise Ration in dBHz (0 ~ 99), empty if not tracking
}

type GPGSV struct {
	Message
	NbOfMessage      int // Number of messages, total number of GPGSV messages being output (1 ~ 3)
	SequenceNumber   int // Sequence number of this entry (1 ~ 3)
	SatellitesInView int
	Satellites       [3]Satellite
}

func (m *GPGSV) GetMessage() *Message { // Implement NMEA interface
	return &m.Message
}

func (m *GPGSV) parse() (err error) {
	if len(m.Fields) != 19 {
		return fmt.Errorf("Incomplete GPGSV message, not enougth data fields (got: %d, wanted: %d)", len(m.Fields), 19)
	}

	if m.NbOfMessage, err = strconv.Atoi(m.Fields[0]); err != nil {
		return
	}

	if m.NbOfMessage < 1 || m.NbOfMessage > 3 {
		return fmt.Errorf("GPGSV number of messages out of range (got: %d)", m.NbOfMessage)
	}

	if m.SequenceNumber, err = strconv.Atoi(m.Fields[1]); err != nil {
		return
	}

	if m.SequenceNumber < 1 || m.SequenceNumber > 3 {
		return fmt.Errorf("GPGSV sequence number out of range (got: %d)", m.SequenceNumber)
	}

	if m.SatellitesInView, err = strconv.Atoi(m.Fields[2]); err != nil {
		return
	}

	return nil
}
