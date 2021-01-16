package nmea

import (
	"fmt"
	"strconv"
	"strings"
)

// Examples:
// $GPGSV,3,1,12,01,05,060,18,02,17,259,43,04,56,287,28,09,08,277,28*77
// $GPGSV,3,2,12,10,34,195,46,13,08,125,45,17,67,014,,20,32,048,24*74
// $GPGSV,3,3,12,23,13,094,48,24,04,292,24,28,49,178,46,32,06,037,22*7D

func NewGPGSV(m Message) *GPGSV {
	return &GPGSV{Message: m}
}

type Satellite struct {
	ID        string
	Elevation *int // Elevation in degree (0 ~ 90)
	Azimuth   *int // Azimuth in degree (0 ~ 359)
	SNR       *int // Signal to Noise Ration in dBHz (0 ~ 99), empty if not tracking
}

func newSatelliteFromFields(f []string) (s Satellite, err error) {
	if len(f) < 4 {
		return s, fmt.Errorf("Not enought fields for create satellite")
	}

	s.ID = f[0]

	if el := strings.TrimSpace(f[1]); len(el) > 0 {
		var elevation int
		elevation, err = strconv.Atoi(el)
		if err != nil {
			return
		}
		s.Elevation = &elevation
	}

	if az := strings.TrimSpace(f[2]); len(az) > 0 {
		var azimuth int
		azimuth, err = strconv.Atoi(az)
		if err != nil {
			return
		}
		s.Azimuth = &azimuth
	}

	if snrStr := strings.TrimSpace(f[3]); len(snrStr) > 0 {
		var snr int
		if snr, err = strconv.Atoi(snrStr); err != nil {
			return
		}
		s.SNR = &snr
	}

	return
}

type GPGSV struct {
	Message
	NbOfMessage      int // Number of messages, total number of GPGSV messages being output (1 ~ 3)
	SequenceNumber   int // Sequence number of this entry (1 ~ 3)
	SatellitesInView int
	Satellites       []Satellite
}

func (m *GPGSV) parse() (err error) {
	if len(m.Fields) < 3 || (len(m.Fields)-3)%4 != 0 {
		return m.Error(fmt.Errorf("Invalid message size (got: %d)", len(m.Fields)))
	}

	if m.NbOfMessage, err = strconv.Atoi(m.Fields[0]); err != nil {
		return m.Error(err)
	}

	if m.NbOfMessage < 1 || m.NbOfMessage > 3 {
		return m.Error(fmt.Errorf("Number of messages out of range (got: %d)", m.NbOfMessage))
	}

	if m.SequenceNumber, err = strconv.Atoi(m.Fields[1]); err != nil {
		return m.Error(err)
	}

	if m.SequenceNumber < 1 || m.SequenceNumber > 3 {
		return m.Error(fmt.Errorf("Sequence number out of range (got: %d)", m.SequenceNumber))
	}

	if m.SatellitesInView, err = strconv.Atoi(m.Fields[2]); err != nil {
		return m.Error(err)
	}

	if m.SatellitesInView > 0 {
		offset := 3
		padding := 4
		m.Satellites = make([]Satellite, 0)

		for len(m.Fields[offset:]) != 0 {
			if len(m.Fields[offset:]) < padding {
				return m.Error(fmt.Errorf("Wrong number of satellite data (got: %d)", len(m.Fields[offset:])))
			}

			sat, err := newSatelliteFromFields(m.Fields[offset : offset+padding])
			if err != nil {
				return m.Error(err)
			}

			m.Satellites = append(m.Satellites, sat)
			offset += padding

		}

		if len(m.Satellites) > 4 {
			return m.Error(fmt.Errorf("Too much satellite data in this message (got: %d)", len(m.Satellites)))
		}
	}

	return nil
}

func (m GPGSV) Serialize() string { // Implement NMEA interface
	hdr := TypeIDs["GPGSV"]
	fields := make([]string, 0)

	fields = append(fields,
		strconv.Itoa(m.NbOfMessage),
		strconv.Itoa(m.SequenceNumber),
		PrependXZero(float64(m.SatellitesInView), "%.0f", 2),
	)

	for _, s := range m.Satellites {
		fields = append(fields, s.ID)

		if s.Elevation != nil {
			fields = append(fields, PrependXZero(float64(*s.Elevation), "%.0f", 2))
		} else {
			fields = append(fields, "")
		}

		if s.Azimuth != nil {
			fields = append(fields, PrependXZero(float64(*s.Azimuth), "%.0f", 3))
		} else {
			fields = append(fields, "")
		}

		if s.SNR == nil {
			fields = append(fields, "")
		} else {
			fields = append(fields, strconv.Itoa(*s.SNR))
		}
	}

	msg := Message{Type: hdr, Fields: fields}
	msg.Checksum = msg.ComputeChecksum()

	return msg.Serialize()
}
