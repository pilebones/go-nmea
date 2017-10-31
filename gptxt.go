package nmea

import (
	"fmt"
	"strconv"
	"strings"
)

// Examples:
// $GPTXT,01,01,02,ANTSTATUS=OK*3B

func NewGPTXT(m Message) *GPTXT {
	return &GPTXT{Message: m}
}

type GPTXT struct {
	Message

	TotalNbMsgInTx int // Total number of messages in this transmission. (01~99)
	MsgNumInTx     int // Message number in this transmission. (01~99)
	Severity       Severity

	/*
		L80 module supports automatic antenna switching function.
		The GPTXT sentence can be used to identify the status of external active antenna. The status of external active antenna is listed as below:
		1. If ANTSTATUS=OK, it means external active antenna is connected and the module will use external active antenna.
		2. If ANTSTATUS=OPEN, it means open-circuit state is dectected and the internal antenna is used at this time.
		3. If ANTSTATUS=SHORT, it means short circuit state is dectected and the internal antenna is used.
	*/
	TxtMsg string
}

func (m *GPTXT) GetMessage() *Message { // Implement NMEA interface
	return &m.Message
}

func (m *GPTXT) parse() (err error) {
	if len(m.Fields) != 4 {
		return m.Error(fmt.Errorf("Incomplete GPTXT message, not enougth data fields (got: %d, wanted: %d)", len(m.Fields), 4))
	}

	if m.TotalNbMsgInTx, err = strconv.Atoi(m.Fields[0]); err != nil {
		return m.Error(fmt.Errorf("Unable to parse total number of messages in this transmission from data field (got: %s)", m.Fields[0]))
	}

	if m.MsgNumInTx, err = strconv.Atoi(m.Fields[1]); err != nil {
		return m.Error(fmt.Errorf("Unable to parse message number in this transmission from data field (got: %s)", m.Fields[1]))
	}

	if m.Severity, err = ParseSeverity(m.Fields[2]); err != nil {
		return m.Error(fmt.Errorf("Unable to parse message severity from data field (got: %s)", m.Fields[2]))
	}

	m.TxtMsg = strings.Join(m.Fields[3:], " ")

	return nil
}

func (m GPTXT) Serialize() string { // Implement NMEA interface

	hdr := TypeIds["GPTXT"]
	fields := make([]string, 0)

	if m.TotalNbMsgInTx < 10 {
		fields = append(fields, fmt.Sprintf("0%d", m.TotalNbMsgInTx))
	} else {
		fields = append(fields, fmt.Sprintf("%d", m.TotalNbMsgInTx))
	}

	if m.MsgNumInTx < 10 {
		fields = append(fields, fmt.Sprintf("0%d", m.MsgNumInTx))
	} else {
		fields = append(fields, fmt.Sprintf("%d", m.MsgNumInTx))
	}

	fields = append(fields, m.Severity.Serialize(), m.TxtMsg)

	msg := Message{Type: hdr, Fields: fields}
	msg.Checksum = msg.ComputeChecksum()

	return msg.Serialize()
}
