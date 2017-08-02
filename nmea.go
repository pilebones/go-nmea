package nmea

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// NMEA Special chars
	PREFIX          = "$"
	FIELD_DELIMITER = ","
	SUFFIX          = "*"

	// Talker IDs
	TALKER_ID_PROPRIETARY TalkerId = "P"  // P for pro proprietary message
	TALKER_ID_GPS         TalkerId = "GP" // Global Positioning System receiver
	TALKER_ID_LC          TalkerId = "LC" // Loran-C receiver
	TALKER_ID_II          TalkerId = "II" // Integrated Instrumentation
	TALKER_ID_IN          TalkerId = "IN" // Integrated Navigation
	TALKER_ID_EC          TalkerId = "EC" // Electronic Chart Display & Information System (ECDIS)
	TALKER_ID_CD          TalkerId = "CD" // Digital Selective Calling (DSC)
	TALKER_ID_GA          TalkerId = "GA" // Galileo Positioning System
	TALKER_ID_GL          TalkerId = "GL" // GLONASS, according to IEIC 61162-1
	TALKER_ID_GN          TalkerId = "GN" // Mixed GPS and GLONASS data, according to IEIC 61162-1
	TALKER_ID_GB          TalkerId = "GB" // BeiDou (China)
	TALKER_ID_BD          TalkerId = "BD" // BeiDou (China)
	TALKER_ID_QZ          TalkerId = "QZ" // QZSS regional GPS augmentation system (Japan)
)

var TypeIds map[string]TypeId

func init() {
	TypeIds = map[string]TypeId{
		"GPAAM": TypeId{Talker: TALKER_ID_GPS, Code: "AAM"}, // Waypoint Arrival Alarm
		"GPALM": TypeId{Talker: TALKER_ID_GPS, Code: "ALM"}, // GPS Almanac Data
		"GPAPA": TypeId{Talker: TALKER_ID_GPS, Code: "APA"}, // Autopilot Sentence "A"
		"GPAPB": TypeId{Talker: TALKER_ID_GPS, Code: "APB"}, // Autopilot Sentence "B"
		"GPASD": TypeId{Talker: TALKER_ID_GPS, Code: "ASD"}, // Autopilot System Data
		"GPBEC": TypeId{Talker: TALKER_ID_GPS, Code: "BEC"}, // Bearing & Distance to Waypoint, Dead Reckoning
		"GPBOD": TypeId{Talker: TALKER_ID_GPS, Code: "BOD"}, // Bearing, Origin to Destination
		"GPBWC": TypeId{Talker: TALKER_ID_GPS, Code: "BWC"}, // Bearing & Distance to Waypoint, Great Circle
		"GPBWR": TypeId{Talker: TALKER_ID_GPS, Code: "BWR"}, // Bearing & Distance to Waypoint, Rhumb Line
		"GPBWW": TypeId{Talker: TALKER_ID_GPS, Code: "BWW"}, // Bearing, Waypoint to Waypoint
		"GPDBT": TypeId{Talker: TALKER_ID_GPS, Code: "DBT"}, // Depth Below Transducer
		"GPDCN": TypeId{Talker: TALKER_ID_GPS, Code: "DCN"}, // Decca Position
		"GPDPT": TypeId{Talker: TALKER_ID_GPS, Code: "DPT"}, // Depth
		"GPFSI": TypeId{Talker: TALKER_ID_GPS, Code: "FSI"}, // Frequency Set Information
		"GPGGA": TypeId{Talker: TALKER_ID_GPS, Code: "GGA"}, // Global Positioning System Fix Data
		"GPGLC": TypeId{Talker: TALKER_ID_GPS, Code: "GLC"}, // Geographic Position, Loran-C
		"GPGLL": TypeId{Talker: TALKER_ID_GPS, Code: "GLL"}, // Geographic Position, Latitude/Longitude
		"GPGSA": TypeId{Talker: TALKER_ID_GPS, Code: "GSA"}, // GPS DOP and Active Satellites
		"GPGSV": TypeId{Talker: TALKER_ID_GPS, Code: "GSV"}, // GPS Satellites in View
		"GPGXA": TypeId{Talker: TALKER_ID_GPS, Code: "GXA"}, // TRANSIT Position
		"GPHDG": TypeId{Talker: TALKER_ID_GPS, Code: "HDG"}, // Heading, Deviation & Variation
		"GPHDT": TypeId{Talker: TALKER_ID_GPS, Code: "HDT"}, // Heading, True
		"GPHSC": TypeId{Talker: TALKER_ID_GPS, Code: "HSC"}, // Heading Steering Command
		"GPLCD": TypeId{Talker: TALKER_ID_GPS, Code: "LCD"}, // Loran-C Signal Data
		"GPMTA": TypeId{Talker: TALKER_ID_GPS, Code: "MTA"}, // Air Temperature (to be phased out)
		"GPMTW": TypeId{Talker: TALKER_ID_GPS, Code: "MTW"}, // Water Temperature
		"GPMWD": TypeId{Talker: TALKER_ID_GPS, Code: "MWD"}, // Wind Direction
		"GPMWV": TypeId{Talker: TALKER_ID_GPS, Code: "MWV"}, // Wind Speed and Angle
		"GPOLN": TypeId{Talker: TALKER_ID_GPS, Code: "OLN"}, // Omega Lane Numbers
		"GPOSD": TypeId{Talker: TALKER_ID_GPS, Code: "OSD"}, // Own Ship Data
		"GPR00": TypeId{Talker: TALKER_ID_GPS, Code: "R00"}, // Waypoint active route (not standard)
		"GPRMA": TypeId{Talker: TALKER_ID_GPS, Code: "RMA"}, // Recommended Minimum Specific Loran-C Data
		"GPRMB": TypeId{Talker: TALKER_ID_GPS, Code: "RMB"}, // Recommended Minimum Navigation Information
		"GPRMC": TypeId{Talker: TALKER_ID_GPS, Code: "RMC"}, // Recommended Minimum Specific GPS/TRANSIT Data
		"GPROT": TypeId{Talker: TALKER_ID_GPS, Code: "ROT"}, // Rate of Turn
		"GPRPM": TypeId{Talker: TALKER_ID_GPS, Code: "RPM"}, // Revolutions
		"GPRSA": TypeId{Talker: TALKER_ID_GPS, Code: "RSA"}, // Rudder Sensor Angle
		"GPRSD": TypeId{Talker: TALKER_ID_GPS, Code: "RSD"}, // RADAR System Data
		"GPRTE": TypeId{Talker: TALKER_ID_GPS, Code: "RTE"}, // Routes
		"GPSFI": TypeId{Talker: TALKER_ID_GPS, Code: "SFI"}, // Scanning Frequency Information
		"GPSTN": TypeId{Talker: TALKER_ID_GPS, Code: "STN"}, // Multiple Data ID
		"GPTRF": TypeId{Talker: TALKER_ID_GPS, Code: "TRF"}, // Transit Fix Data
		"GPTTM": TypeId{Talker: TALKER_ID_GPS, Code: "TTM"}, // Tracked Target Message
		"GPTXT": TypeId{Talker: TALKER_ID_GPS, Code: "TXT"}, // Tracked Status of External Antenna
		"GPVBW": TypeId{Talker: TALKER_ID_GPS, Code: "VBW"}, // Dual Ground/Water Speed
		"GPVDR": TypeId{Talker: TALKER_ID_GPS, Code: "VDR"}, // Set and Drift
		"GPVHW": TypeId{Talker: TALKER_ID_GPS, Code: "VHW"}, // Water Speed and Heading
		"GPVLW": TypeId{Talker: TALKER_ID_GPS, Code: "VLW"}, // Distance Traveled through the Water
		"GPVPW": TypeId{Talker: TALKER_ID_GPS, Code: "VPW"}, // Speed, Measured Parallel to Wind
		"GPVTG": TypeId{Talker: TALKER_ID_GPS, Code: "VTG"}, // Track Made Good and Ground Speed
		"GPWCV": TypeId{Talker: TALKER_ID_GPS, Code: "WCV"}, // Waypoint Closure Velocity
		"GPWNC": TypeId{Talker: TALKER_ID_GPS, Code: "WNC"}, // Distance, Waypoint to Waypoint
		"GPWPL": TypeId{Talker: TALKER_ID_GPS, Code: "WPL"}, // Waypoint Location
		"GPXDR": TypeId{Talker: TALKER_ID_GPS, Code: "XDR"}, // Transducer Measurements
		"GPXTE": TypeId{Talker: TALKER_ID_GPS, Code: "XTE"}, // Cross-Track Error, Measured
		"GPXTR": TypeId{Talker: TALKER_ID_GPS, Code: "XTR"}, // Cross-Track Error, Dead Reckoning
		"GPZDA": TypeId{Talker: TALKER_ID_GPS, Code: "ZDA"}, // Time & Date
		"GPZFO": TypeId{Talker: TALKER_ID_GPS, Code: "ZFO"}, // UTC & Time from Origin Waypoint
		"GPZTG": TypeId{Talker: TALKER_ID_GPS, Code: "ZTG"}, // UTC & Time to Destination Waypoint
	}
}

type TypeId struct {
	Talker TalkerId
	Code   string
}

func (t TypeId) String() string {
	return t.Talker.String() + t.Code
}

type TalkerId string

func (t TalkerId) String() string {
	return string(t)
}

type Message struct {
	Type     TypeId
	Fields   []string
	Checksum uint8
}

func (m Message) String() string {
	return PREFIX + m.Payload() + SUFFIX + fmt.Sprintf("%x", m.Checksum)
}

func (m Message) Payload() string {
	return m.Type.String() + FIELD_DELIMITER + strings.Join(m.Fields, FIELD_DELIMITER)
}

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
	if len(fields) < 2 {
		return fmt.Errorf("Message should countains at least two fields (got: %d)", len(fields))
	}

	typ, ok := TypeIds[fields[0]]
	if !ok {
		return fmt.Errorf("Message should countains a valid type id (got: %s)", fields[0])
	}
	m.Type = typ

	m.Fields = fields[1:]

	checksum, err := strconv.ParseUint(data[checksumOffset:], 16, 8)
	if err != nil {
		return
	}

	if uint8(checksum) != m.ComputeChecksum() {
		return fmt.Errorf("Checksump mismatch (got: 0x%x, wanted: 0x%x)", checksum, m.ComputeChecksum())
	}

	return nil
}

func Parse(raw string) (m *Message, err error) {
	m = &Message{}
	if err = m.parse(raw); err != nil {
		return nil, err
	}
	return
}
