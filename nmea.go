package nmea

import (
	"fmt"
	"strconv"
	"strings"
)

// Interface for each kind of NMEA message
type NMEA interface {
	GetMessage() *Message
}

// Interface for each kind of NMEA header according to TalkerId
type Header interface {
	GetTypeId() TypeId
	String() string
}

type Message struct {
	Type     Header
	Fields   []string
	Checksum uint8
}

func (m *Message) GetMessage() *Message {
	return m
}

// Serialize message to render raw
func (m Message) String() string {
	output := PREFIX + m.Payload() + SUFFIX
	checksum := fmt.Sprintf("%X", m.Checksum)
	if len(checksum) == 1 {
		checksum = "0" + checksum // Padd with 0 if needed
	}
	return output + checksum
}

// Payload return data after $ and before *
func (m Message) Payload() string {
	if f := strings.Join(m.Fields, FIELD_DELIMITER); len(f) > 0 {
		return m.Type.String() + FIELD_DELIMITER + strings.Join(m.Fields, FIELD_DELIMITER)
	}
	return m.Type.String()
}

// ComputeChecksum recompute checksum from extracted payload
func (m Message) ComputeChecksum() (c uint8) {
	for i := 0; i < len(m.Payload()); i++ {
		c ^= m.Payload()[i]
	}
	return
}

func (m *Message) parse(data string) (err error) {
	if len(data) < (len(PREFIX) + len(SUFFIX) + 2) { // +2 for checksum in hex format
		return fmt.Errorf("Wrong length")
	}

	startMsgOffset := 0
	endMsgOffset := len(data) - 3
	checksumOffset := len(data) - 2

	if string(data[startMsgOffset]) != PREFIX {
		return fmt.Errorf("Message should start with %s (got: %s)", PREFIX, string(data[startMsgOffset]))
	}

	if string(data[endMsgOffset]) != SUFFIX {
		return fmt.Errorf("Message should countains with %s (got: %s)", SUFFIX, string(data[endMsgOffset]))
	}

	msg := data[startMsgOffset+1 : endMsgOffset]

	fields := strings.Split(msg, FIELD_DELIMITER)
	if len(fields) == 0 {
		return fmt.Errorf("Message has no type or field")
	}

	typ, ok := TypeIds[fields[0]]
	if !ok {
		return fmt.Errorf("Message should countains a valid type id (got: %s)", fields[0])
	}
	m.Type = typ

	if len(fields) > 1 {
		m.Fields = fields[1:]
	}

	checksum, err := strconv.ParseUint(data[checksumOffset:], 16, 8)
	if err != nil {
		return
	}

	if m.Checksum = uint8(checksum); m.Checksum != m.ComputeChecksum() {
		return fmt.Errorf("Checksump mismatch (got: 0x%x, wanted: 0x%x)", checksum, m.ComputeChecksum())
	}

	return nil
}

// Parse return message for any kind of NMEA message raw
func Parse(raw string) (NMEA, error) {
	var err error
	m := &Message{}
	if err = m.parse(raw); err != nil {
		return nil, err
	}

	switch m.Type.String() {
	case "GPRMC":
		gprmc := NewGPRMC(*m)
		err = gprmc.parse()
		return gprmc, err
	case "GPVTG":
		gpvtg := NewGPVTG(*m)
		err = gpvtg.parse()
		return gpvtg, err
	case "GPGGA":
		gpgga := NewGPGGA(*m)
		err = gpgga.parse()
		return gpgga, err
	case "GPGSA":
		gpgsa := NewGPGSA(*m)
		err = gpgsa.parse()
		return gpgsa, err
	}

	return m, err
}
