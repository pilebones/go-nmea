package nmea

import (
	"fmt"
	"strconv"
	"strings"
)

// NMEA is an interface for each kind of NMEA message
type NMEA interface {
	GetMessage() Message
	Error(err error) error
	Serialize() string
}

// Header is an interface for each kind of NMEA header according to TalkerId
type Header interface {
	GetTypeID() TypeID
	Serialize() string
}

// Message is the base aand low-level struct (envelope) without advanced dissection for each NMEA message
type Message struct {
	Type     Header
	Fields   []string
	Checksum uint8
}

// GetMessage return base Message to respect interface
func (m Message) GetMessage() Message {
	return m
}

// Error return common error with wrapped data to enhance debugging
func (m Message) Error(err error) error {
	return fmt.Errorf("[%s] %s (with payload: %s)", m.Type.Serialize(), err.Error(), strings.Join(m.Fields, FieldDelimiter))
}

// Serialize NMEA message to render raw
func (m Message) Serialize() string {
	output := Prefix + m.Payload() + Suffix
	checksum := fmt.Sprintf("%X", m.Checksum)
	if len(checksum) == 1 {
		checksum = "0" + checksum // Padd with 0 if needed
	}
	return output + checksum
}

// Payload return data after $ and before *
func (m Message) Payload() string {
	if f := strings.Join(m.Fields, FieldDelimiter); len(f) > 0 {
		return m.Type.Serialize() + FieldDelimiter + f
	}
	return m.Type.Serialize()
}

// ComputeChecksum recompute checksum from extracted payload
func (m Message) ComputeChecksum() (c uint8) {
	for i := 0; i < len(m.Payload()); i++ {
		c ^= m.Payload()[i]
	}
	return
}

func (m *Message) parse(data string) (err error) {
	if len(data) < (len(Prefix) + len(Suffix) + 2) { // +2 for checksum in hex format
		return fmt.Errorf("Wrong length")
	}

	startMsgOffset := 0
	endMsgOffset := len(data) - 3
	checksumOffset := len(data) - 2

	if string(data[startMsgOffset]) != Prefix {
		return fmt.Errorf("Message should start with %s (got: %s)", Prefix, string(data[startMsgOffset]))
	}

	if string(data[endMsgOffset]) != Suffix {
		return fmt.Errorf("Message should countains with %s (got: %s)", Suffix, string(data[endMsgOffset]))
	}

	msg := data[startMsgOffset+1 : endMsgOffset]

	fields := strings.Split(msg, FieldDelimiter)
	if len(fields) == 0 {
		return fmt.Errorf("Message has no type or field")
	}

	typ, ok := TypeIDs[fields[0]]
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
		return m.Error(fmt.Errorf("Checksump mismatch (got: 0x%x, wanted: 0x%x)", checksum, m.ComputeChecksum()))
	}

	return nil
}

// Parse return message for any kind of NMEA message raw
func Parse(raw string) (NMEA, error) {
	var err error
	m := &Message{}

	raw = strings.TrimRight(raw, "\n") // Remove residual CRLF chars

	if err = m.parse(raw); err != nil {
		return nil, err
	}

	switch m.Type.Serialize() {
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
	case "GPGSV":
		gpgsv := NewGPGSV(*m)
		err = gpgsv.parse()
		return gpgsv, err
	case "GPGLL":
		gpgll := NewGPGLL(*m)
		err = gpgll.parse()
		return gpgll, err
	case "GPTXT":
		gptxt := NewGPTXT(*m)
		err = gptxt.parse()
		return gptxt, err
	}

	return m, err
}
