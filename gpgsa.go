package nmea

import (
	"fmt"
	"strconv"
)

// Examples:
// $GPGSA,A,3,14,06,16,31,23,,,,,,,,1.66,1.42,0.84*0F

func NewGPGSA(m Message) *GPGSA {
	return &GPGSA{Message: m}
}

type GPGSA struct {
	Message

	Mode      Mode
	FixStatus FixStatus
}

func (m *GPGSA) GetMessage() *Message { // Implement NMEA interface
	return &m.Message
}

func (m *GPGSA) parse() (err error) {
	if len(m.Fields) != 17 {
		return fmt.Errorf("Incomplete GPGSA message, not enougth data fields (got: %d, wanted: %d)", len(m.Fields), 17)
	}

	if m.Mode, err = ParseMode(m.Fields[0]); err != nil {
		return
	}

	if m.FixStatus, err = ParseFixStatus(m.Fields[1]); err != nil {
		return
	}

	return nil
}

const (
	_ = iota
	FIX_STATUS_NO_FIX
	FIX_STATUS_2D
	FIX_STATUS_3D
)

type FixStatus int

func (s FixStatus) String() string {
	switch s {
	case FIX_STATUS_NO_FIX:
		return "No fix"
	case FIX_STATUS_2D:
		return "2D fix"
	case FIX_STATUS_3D:
		return "3D fix"
	default:
		return "unknow"
	}
}

func ParseFixStatus(raw string) (fs FixStatus, err error) {
	i, err := strconv.ParseInt(raw, 10, 0)
	if err != nil {
		return
	}

	fs = FixStatus(i)
	switch fs {
	case FIX_STATUS_NO_FIX, FIX_STATUS_2D, FIX_STATUS_3D:
	default:
		err = fmt.Errorf("unknow value (got: %d)", i)
	}
	return
}

const (
	MODE_MANUAL Mode = "M"
	MODE_AUTO   Mode = "A"
)

type Mode string

func (m Mode) String() string {
	return string(m)
}

func ParseMode(raw string) (m Mode, err error) {
	m = Mode(raw)
	switch m {
	case MODE_MANUAL, MODE_AUTO:
	default:
		err = fmt.Errorf("unknow value (got: %s)", m)
	}
	return
}
