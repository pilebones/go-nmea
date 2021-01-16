package nmea

import "fmt"

const (
	// NMEA special-chars
	// Prefix is special char to begin NMEA message
	Prefix = "$"
	// FieldDelimiter is special char to delimit a field in NMEA message
	FieldDelimiter = ","
	// Suffix is special char to finish NMEA message
	Suffix = "*"

	// Talker IDs
	TalkerIDProprietary TalkerID = "P"  // P for pro proprietary message
	TalkerIDGPS         TalkerID = "GP" // Global Positioning System receiver
	TalkerIDLC          TalkerID = "LC" // Loran-C receiver
	TalkerIDII          TalkerID = "II" // Integrated Instrumentation
	TalkerIDIN          TalkerID = "IN" // Integrated Navigation
	TalkerIDEC          TalkerID = "EC" // Electronic Chart Display & Information System (ECDIS)
	TalkerIDCD          TalkerID = "CD" // Digital Selective Calling (DSC)
	TalkerIDGA          TalkerID = "GA" // Galileo Positioning System
	TalkerIDGL          TalkerID = "GL" // GLONASS, according to IEIC 61162-1
	TalkerIDGN          TalkerID = "GN" // Mixed GPS and GLONASS data, according to IEIC 61162-1
	TalkerIDGB          TalkerID = "GB" // BeiDou (China)
	TalkerIDBD          TalkerID = "BD" // BeiDou (China)
	TalkerIDQZ          TalkerID = "QZ" // QZSS regional GPS augmentation system (Japan)
)

type TypeID struct {
	Talker TalkerID
	Code   string
}

func (t TypeID) GetTypeID() TypeID {
	return t
}

func (t TypeID) Serialize() string {
	return t.Talker.Serialize() + t.Code
}

type MtkTypeID struct {
	TypeID
	PacketType string
}

func (t MtkTypeID) Serialize() string {
	return t.TypeID.Serialize() + t.PacketType
}

type TalkerID string

func (t TalkerID) Serialize() string {
	return string(t)
}

// TypeIDs is a dictionnary of all kind of NMEA message header by full-code
var TypeIDs map[string]Header

func init() {
	TypeIDs = map[string]Header{
		"GPAAM":   TypeID{Talker: TalkerIDGPS, Code: "AAM"},                                               // Waypoint Arrival Alarm
		"GPALM":   TypeID{Talker: TalkerIDGPS, Code: "ALM"},                                               // GPS Almanac Data
		"GPAPA":   TypeID{Talker: TalkerIDGPS, Code: "APA"},                                               // Autopilot Sentence "A"
		"GPAPB":   TypeID{Talker: TalkerIDGPS, Code: "APB"},                                               // Autopilot Sentence "B"
		"GPASD":   TypeID{Talker: TalkerIDGPS, Code: "ASD"},                                               // Autopilot System Data
		"GPBEC":   TypeID{Talker: TalkerIDGPS, Code: "BEC"},                                               // Bearing & Distance to Waypoint, Dead Reckoning
		"GPBOD":   TypeID{Talker: TalkerIDGPS, Code: "BOD"},                                               // Bearing, Origin to Destination
		"GPBWC":   TypeID{Talker: TalkerIDGPS, Code: "BWC"},                                               // Bearing & Distance to Waypoint, Great Circle
		"GPBWR":   TypeID{Talker: TalkerIDGPS, Code: "BWR"},                                               // Bearing & Distance to Waypoint, Rhumb Line
		"GPBWW":   TypeID{Talker: TalkerIDGPS, Code: "BWW"},                                               // Bearing, Waypoint to Waypoint
		"GPDBT":   TypeID{Talker: TalkerIDGPS, Code: "DBT"},                                               // Depth Below Transducer
		"GPDCN":   TypeID{Talker: TalkerIDGPS, Code: "DCN"},                                               // Decca Position
		"GPDPT":   TypeID{Talker: TalkerIDGPS, Code: "DPT"},                                               // Depth
		"GPFSI":   TypeID{Talker: TalkerIDGPS, Code: "FSI"},                                               // Frequency Set Information
		"GPGGA":   TypeID{Talker: TalkerIDGPS, Code: "GGA"},                                               // Global Positioning System Fix Data
		"GPGLC":   TypeID{Talker: TalkerIDGPS, Code: "GLC"},                                               // Geographic Position, Loran-C
		"GPGLL":   TypeID{Talker: TalkerIDGPS, Code: "GLL"},                                               // Geographic Position, Latitude/Longitude
		"GPGSA":   TypeID{Talker: TalkerIDGPS, Code: "GSA"},                                               // GPS DOP and Active Satellites
		"GPGSV":   TypeID{Talker: TalkerIDGPS, Code: "GSV"},                                               // GPS Satellites in View
		"GPGXA":   TypeID{Talker: TalkerIDGPS, Code: "GXA"},                                               // TRANSIT Position
		"GPHDG":   TypeID{Talker: TalkerIDGPS, Code: "HDG"},                                               // Heading, Deviation & Variation
		"GPHDT":   TypeID{Talker: TalkerIDGPS, Code: "HDT"},                                               // Heading, True
		"GPHSC":   TypeID{Talker: TalkerIDGPS, Code: "HSC"},                                               // Heading Steering Command
		"GPLCD":   TypeID{Talker: TalkerIDGPS, Code: "LCD"},                                               // Loran-C Signal Data
		"GPMTA":   TypeID{Talker: TalkerIDGPS, Code: "MTA"},                                               // Air Temperature (to be phased out)
		"GPMTW":   TypeID{Talker: TalkerIDGPS, Code: "MTW"},                                               // Water Temperature
		"GPMWD":   TypeID{Talker: TalkerIDGPS, Code: "MWD"},                                               // Wind Direction
		"GPMWV":   TypeID{Talker: TalkerIDGPS, Code: "MWV"},                                               // Wind Speed and Angle
		"GPOLN":   TypeID{Talker: TalkerIDGPS, Code: "OLN"},                                               // Omega Lane Numbers
		"GPOSD":   TypeID{Talker: TalkerIDGPS, Code: "OSD"},                                               // Own Ship Data
		"GPR00":   TypeID{Talker: TalkerIDGPS, Code: "R00"},                                               // Waypoint active route (not standard)
		"GPRMA":   TypeID{Talker: TalkerIDGPS, Code: "RMA"},                                               // Recommended Minimum Specific Loran-C Data
		"GPRMB":   TypeID{Talker: TalkerIDGPS, Code: "RMB"},                                               // Recommended Minimum Navigation Information
		"GPRMC":   TypeID{Talker: TalkerIDGPS, Code: "RMC"},                                               // Recommended Minimum Specific GPS/TRANSIT Data
		"GPROT":   TypeID{Talker: TalkerIDGPS, Code: "ROT"},                                               // Rate of Turn
		"GPRPM":   TypeID{Talker: TalkerIDGPS, Code: "RPM"},                                               // Revolutions
		"GPRSA":   TypeID{Talker: TalkerIDGPS, Code: "RSA"},                                               // Rudder Sensor Angle
		"GPRSD":   TypeID{Talker: TalkerIDGPS, Code: "RSD"},                                               // RADAR System Data
		"GPRTE":   TypeID{Talker: TalkerIDGPS, Code: "RTE"},                                               // Routes
		"GPSFI":   TypeID{Talker: TalkerIDGPS, Code: "SFI"},                                               // Scanning Frequency Information
		"GPSTN":   TypeID{Talker: TalkerIDGPS, Code: "STN"},                                               // Multiple Data ID
		"GPTRF":   TypeID{Talker: TalkerIDGPS, Code: "TRF"},                                               // Transit Fix Data
		"GPTTM":   TypeID{Talker: TalkerIDGPS, Code: "TTM"},                                               // Tracked Target Message
		"GPTXT":   TypeID{Talker: TalkerIDGPS, Code: "TXT"},                                               // Tracked Status of External Antenna
		"GPVBW":   TypeID{Talker: TalkerIDGPS, Code: "VBW"},                                               // Dual Ground/Water Speed
		"GPVDR":   TypeID{Talker: TalkerIDGPS, Code: "VDR"},                                               // Set and Drift
		"GPVHW":   TypeID{Talker: TalkerIDGPS, Code: "VHW"},                                               // Water Speed and Heading
		"GPVLW":   TypeID{Talker: TalkerIDGPS, Code: "VLW"},                                               // Distance Traveled through the Water
		"GPVPW":   TypeID{Talker: TalkerIDGPS, Code: "VPW"},                                               // Speed, Measured Parallel to Wind
		"GPVTG":   TypeID{Talker: TalkerIDGPS, Code: "VTG"},                                               // Track Made Good and Ground Speed
		"GPWCV":   TypeID{Talker: TalkerIDGPS, Code: "WCV"},                                               // Waypoint Closure Velocity
		"GPWNC":   TypeID{Talker: TalkerIDGPS, Code: "WNC"},                                               // Distance, Waypoint to Waypoint
		"GPWPL":   TypeID{Talker: TalkerIDGPS, Code: "WPL"},                                               // Waypoint Location
		"GPXDR":   TypeID{Talker: TalkerIDGPS, Code: "XDR"},                                               // Transducer Measurements
		"GPXTE":   TypeID{Talker: TalkerIDGPS, Code: "XTE"},                                               // Cross-Track Error, Measured
		"GPXTR":   TypeID{Talker: TalkerIDGPS, Code: "XTR"},                                               // Cross-Track Error, Dead Reckoning
		"GPZDA":   TypeID{Talker: TalkerIDGPS, Code: "ZDA"},                                               // Time & Date
		"GPZFO":   TypeID{Talker: TalkerIDGPS, Code: "ZFO"},                                               // UTC & Time from Origin Waypoint
		"GPZTG":   TypeID{Talker: TalkerIDGPS, Code: "ZTG"},                                               // UTC & Time to Destination Waypoint
		"PMTK010": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "010"}, // PMTK_SYS_MSG
		"PMTK011": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "011"}, // PMTK_TXT_MSG
		"PMTK001": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "001"}, // PMTK_ACK
		"PMTK101": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "101"}, // PMTK_CMD_HOT_START
		"PMTK102": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "102"}, // PMTK_CMD_WARM_START
		"PMTK103": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "103"}, // PMTK_CMD_COLD_START
		"PMTK104": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "104"}, // PMTK_CMD_FULL_COLD_START
		"PMTK161": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "161"}, // PMTK_CMD_STANDBY_MODE
		"PMTK183": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "183"}, // PMTK_LOCUS_QUERY_STATUS
		"PMTKLOG": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "LOG"}, // PMTK_LOG
		"PMTK184": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "184"}, // PMTK_LOCUS_ERASE_FLASH
		"PMTK185": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "185"}, // PMTK_LOCUS_STOP_LOGGER
		"PMTK622": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "622"}, // PMTK_Q_LOCUS_DATA
		"PMTK225": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "225"}, // PMTK_SET_PERIODIC
		"PMTK251": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "251"}, // PMTK_SET_NMEA_BAUDRATE
		"PMTK286": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "286"}, // PMTK_SET_AIC_ENABLED
		"PMTK300": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "300"}, // PMTK_API_SET_FIX_CTL
		"PMTK301": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "301"}, // PMTK_API_SET_DGPS_MODE
		"PMTK313": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "313"}, // PMTK_API_SET_SBAS_ENABLED
		"PMTK314": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "314"}, // PMTK_API_SET_NMEA_OUTPUT
		"PMTK386": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "386"}, // PMTK_API_SET_STATIC_NAV_THD
		"PMTK400": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "400"}, // PMTK_API_Q_FIX_CTL
		"PMTK401": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "401"}, // PMTK_API_Q_DGPS_MODE
		"PMTK413": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "413"}, // PMTK_API_Q_SBAS_ENABLED
		"PMTK414": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "414"}, // PMTK_API_Q_NMEA_OUTPUT
		"PMTK605": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "605"}, // PMTK_Q_RELEASE
		"PMTK500": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "500"}, // PMTK_DT_FIX_CTL
		"PMTK501": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "501"}, // PMTK_DT_DGPS_MODE
		"PMTK513": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "513"}, // PMTK_DT_SBAS_ENABLED
		"PMTK514": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "514"}, // PMTK_DT_NMEA_OUTPUT
		"PMTK705": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "705"}, // PMTK_DT_RELEASE
		"PMTK869": MtkTypeID{TypeID: TypeID{Talker: TalkerIDProprietary, Code: "MTK"}, PacketType: "869"}, // PMTK_EASY_ENABLE
	}
}

const (
	Valid   DataValid = true
	Invalid DataValid = false
)

type DataValid bool

func (v DataValid) Serialize() string {
	if v == Valid {
		return "A"
	}
	return "V"
}

const (
	NoFixMode           PositioningMode = "N"
	AutonomousGNSSFix   PositioningMode = "A"
	DifferentialGNSSFix PositioningMode = "D"
)

type PositioningMode string

func (p PositioningMode) Serialize() string {
	return string(p)
}

func (p PositioningMode) String() string {
	switch p {
	case NoFixMode:
		return "No fix"
	case AutonomousGNSSFix:
		return "Autonomous GNSS fix"
	case DifferentialGNSSFix:
		return "Differential GNSS fix"
	default:
		return "unknow"
	}
}

func ParsePositioningMode(raw string) (pm PositioningMode, err error) {
	pm = PositioningMode(raw)
	switch pm {
	case NoFixMode, AutonomousGNSSFix, DifferentialGNSSFix:
	default:
		err = fmt.Errorf("unknow value")
	}
	return
}

const (
	ERROR   Severity = "00"
	WARNING Severity = "01"
	NOTICE  Severity = "02"
	USER    Severity = "07"
)

type Severity string

func (s Severity) Serialize() string {
	return string(s)
}

func (s Severity) String() string {
	switch s {
	case ERROR:
		return "ERROR"
	case WARNING:
		return "WARNING"
	case NOTICE:
		return "NOTICE"
	case USER:
		return "USER"
	default:
		return "unknow"
	}
}

func ParseSeverity(raw string) (s Severity, err error) {
	s = Severity(raw)
	switch s {
	case ERROR, WARNING, NOTICE, USER:
	default:
		err = fmt.Errorf("unknow value")
	}
	return
}
