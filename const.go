package nmea

import "fmt"

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

type TypeId struct {
	Talker TalkerId
	Code   string
}

func (t TypeId) GetTypeId() TypeId {
	return t
}

func (t TypeId) String() string {
	return t.Talker.String() + t.Code
}

type MtkTypeId struct {
	TypeId
	PacketType string
}

func (t MtkTypeId) String() string {
	return t.TypeId.String() + t.PacketType
}

type TalkerId string

func (t TalkerId) String() string {
	return string(t)
}

// Dictionnary of all kind of NMEA message header by full-code
var TypeIds map[string]Header

func init() {
	TypeIds = map[string]Header{
		"GPAAM":   TypeId{Talker: TALKER_ID_GPS, Code: "AAM"},                                               // Waypoint Arrival Alarm
		"GPALM":   TypeId{Talker: TALKER_ID_GPS, Code: "ALM"},                                               // GPS Almanac Data
		"GPAPA":   TypeId{Talker: TALKER_ID_GPS, Code: "APA"},                                               // Autopilot Sentence "A"
		"GPAPB":   TypeId{Talker: TALKER_ID_GPS, Code: "APB"},                                               // Autopilot Sentence "B"
		"GPASD":   TypeId{Talker: TALKER_ID_GPS, Code: "ASD"},                                               // Autopilot System Data
		"GPBEC":   TypeId{Talker: TALKER_ID_GPS, Code: "BEC"},                                               // Bearing & Distance to Waypoint, Dead Reckoning
		"GPBOD":   TypeId{Talker: TALKER_ID_GPS, Code: "BOD"},                                               // Bearing, Origin to Destination
		"GPBWC":   TypeId{Talker: TALKER_ID_GPS, Code: "BWC"},                                               // Bearing & Distance to Waypoint, Great Circle
		"GPBWR":   TypeId{Talker: TALKER_ID_GPS, Code: "BWR"},                                               // Bearing & Distance to Waypoint, Rhumb Line
		"GPBWW":   TypeId{Talker: TALKER_ID_GPS, Code: "BWW"},                                               // Bearing, Waypoint to Waypoint
		"GPDBT":   TypeId{Talker: TALKER_ID_GPS, Code: "DBT"},                                               // Depth Below Transducer
		"GPDCN":   TypeId{Talker: TALKER_ID_GPS, Code: "DCN"},                                               // Decca Position
		"GPDPT":   TypeId{Talker: TALKER_ID_GPS, Code: "DPT"},                                               // Depth
		"GPFSI":   TypeId{Talker: TALKER_ID_GPS, Code: "FSI"},                                               // Frequency Set Information
		"GPGGA":   TypeId{Talker: TALKER_ID_GPS, Code: "GGA"},                                               // Global Positioning System Fix Data
		"GPGLC":   TypeId{Talker: TALKER_ID_GPS, Code: "GLC"},                                               // Geographic Position, Loran-C
		"GPGLL":   TypeId{Talker: TALKER_ID_GPS, Code: "GLL"},                                               // Geographic Position, Latitude/Longitude
		"GPGSA":   TypeId{Talker: TALKER_ID_GPS, Code: "GSA"},                                               // GPS DOP and Active Satellites
		"GPGSV":   TypeId{Talker: TALKER_ID_GPS, Code: "GSV"},                                               // GPS Satellites in View
		"GPGXA":   TypeId{Talker: TALKER_ID_GPS, Code: "GXA"},                                               // TRANSIT Position
		"GPHDG":   TypeId{Talker: TALKER_ID_GPS, Code: "HDG"},                                               // Heading, Deviation & Variation
		"GPHDT":   TypeId{Talker: TALKER_ID_GPS, Code: "HDT"},                                               // Heading, True
		"GPHSC":   TypeId{Talker: TALKER_ID_GPS, Code: "HSC"},                                               // Heading Steering Command
		"GPLCD":   TypeId{Talker: TALKER_ID_GPS, Code: "LCD"},                                               // Loran-C Signal Data
		"GPMTA":   TypeId{Talker: TALKER_ID_GPS, Code: "MTA"},                                               // Air Temperature (to be phased out)
		"GPMTW":   TypeId{Talker: TALKER_ID_GPS, Code: "MTW"},                                               // Water Temperature
		"GPMWD":   TypeId{Talker: TALKER_ID_GPS, Code: "MWD"},                                               // Wind Direction
		"GPMWV":   TypeId{Talker: TALKER_ID_GPS, Code: "MWV"},                                               // Wind Speed and Angle
		"GPOLN":   TypeId{Talker: TALKER_ID_GPS, Code: "OLN"},                                               // Omega Lane Numbers
		"GPOSD":   TypeId{Talker: TALKER_ID_GPS, Code: "OSD"},                                               // Own Ship Data
		"GPR00":   TypeId{Talker: TALKER_ID_GPS, Code: "R00"},                                               // Waypoint active route (not standard)
		"GPRMA":   TypeId{Talker: TALKER_ID_GPS, Code: "RMA"},                                               // Recommended Minimum Specific Loran-C Data
		"GPRMB":   TypeId{Talker: TALKER_ID_GPS, Code: "RMB"},                                               // Recommended Minimum Navigation Information
		"GPRMC":   TypeId{Talker: TALKER_ID_GPS, Code: "RMC"},                                               // Recommended Minimum Specific GPS/TRANSIT Data
		"GPROT":   TypeId{Talker: TALKER_ID_GPS, Code: "ROT"},                                               // Rate of Turn
		"GPRPM":   TypeId{Talker: TALKER_ID_GPS, Code: "RPM"},                                               // Revolutions
		"GPRSA":   TypeId{Talker: TALKER_ID_GPS, Code: "RSA"},                                               // Rudder Sensor Angle
		"GPRSD":   TypeId{Talker: TALKER_ID_GPS, Code: "RSD"},                                               // RADAR System Data
		"GPRTE":   TypeId{Talker: TALKER_ID_GPS, Code: "RTE"},                                               // Routes
		"GPSFI":   TypeId{Talker: TALKER_ID_GPS, Code: "SFI"},                                               // Scanning Frequency Information
		"GPSTN":   TypeId{Talker: TALKER_ID_GPS, Code: "STN"},                                               // Multiple Data ID
		"GPTRF":   TypeId{Talker: TALKER_ID_GPS, Code: "TRF"},                                               // Transit Fix Data
		"GPTTM":   TypeId{Talker: TALKER_ID_GPS, Code: "TTM"},                                               // Tracked Target Message
		"GPTXT":   TypeId{Talker: TALKER_ID_GPS, Code: "TXT"},                                               // Tracked Status of External Antenna
		"GPVBW":   TypeId{Talker: TALKER_ID_GPS, Code: "VBW"},                                               // Dual Ground/Water Speed
		"GPVDR":   TypeId{Talker: TALKER_ID_GPS, Code: "VDR"},                                               // Set and Drift
		"GPVHW":   TypeId{Talker: TALKER_ID_GPS, Code: "VHW"},                                               // Water Speed and Heading
		"GPVLW":   TypeId{Talker: TALKER_ID_GPS, Code: "VLW"},                                               // Distance Traveled through the Water
		"GPVPW":   TypeId{Talker: TALKER_ID_GPS, Code: "VPW"},                                               // Speed, Measured Parallel to Wind
		"GPVTG":   TypeId{Talker: TALKER_ID_GPS, Code: "VTG"},                                               // Track Made Good and Ground Speed
		"GPWCV":   TypeId{Talker: TALKER_ID_GPS, Code: "WCV"},                                               // Waypoint Closure Velocity
		"GPWNC":   TypeId{Talker: TALKER_ID_GPS, Code: "WNC"},                                               // Distance, Waypoint to Waypoint
		"GPWPL":   TypeId{Talker: TALKER_ID_GPS, Code: "WPL"},                                               // Waypoint Location
		"GPXDR":   TypeId{Talker: TALKER_ID_GPS, Code: "XDR"},                                               // Transducer Measurements
		"GPXTE":   TypeId{Talker: TALKER_ID_GPS, Code: "XTE"},                                               // Cross-Track Error, Measured
		"GPXTR":   TypeId{Talker: TALKER_ID_GPS, Code: "XTR"},                                               // Cross-Track Error, Dead Reckoning
		"GPZDA":   TypeId{Talker: TALKER_ID_GPS, Code: "ZDA"},                                               // Time & Date
		"GPZFO":   TypeId{Talker: TALKER_ID_GPS, Code: "ZFO"},                                               // UTC & Time from Origin Waypoint
		"GPZTG":   TypeId{Talker: TALKER_ID_GPS, Code: "ZTG"},                                               // UTC & Time to Destination Waypoint
		"PMTK010": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "010"}, // PMTK_SYS_MSG
		"PMTK011": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "011"}, // PMTK_TXT_MSG
		"PMTK001": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "001"}, // PMTK_ACK
		"PMTK101": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "101"}, // PMTK_CMD_HOT_START
		"PMTK102": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "102"}, // PMTK_CMD_WARM_START
		"PMTK103": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "103"}, // PMTK_CMD_COLD_START
		"PMTK104": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "104"}, // PMTK_CMD_FULL_COLD_START
		"PMTK161": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "161"}, // PMTK_CMD_STANDBY_MODE
		"PMTK183": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "183"}, // PMTK_LOCUS_QUERY_STATUS
		"PMTKLOG": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "LOG"}, // PMTK_LOG
		"PMTK184": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "184"}, // PMTK_LOCUS_ERASE_FLASH
		"PMTK185": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "185"}, // PMTK_LOCUS_STOP_LOGGER
		"PMTK622": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "622"}, // PMTK_Q_LOCUS_DATA
		"PMTK225": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "225"}, // PMTK_SET_PERIODIC
		"PMTK251": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "251"}, // PMTK_SET_NMEA_BAUDRATE
		"PMTK286": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "286"}, // PMTK_SET_AIC_ENABLED
		"PMTK300": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "300"}, // PMTK_API_SET_FIX_CTL
		"PMTK301": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "301"}, // PMTK_API_SET_DGPS_MODE
		"PMTK313": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "313"}, // PMTK_API_SET_SBAS_ENABLED
		"PMTK314": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "314"}, // PMTK_API_SET_NMEA_OUTPUT
		"PMTK386": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "386"}, // PMTK_API_SET_STATIC_NAV_THD
		"PMTK400": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "400"}, // PMTK_API_Q_FIX_CTL
		"PMTK401": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "401"}, // PMTK_API_Q_DGPS_MODE
		"PMTK413": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "413"}, // PMTK_API_Q_SBAS_ENABLED
		"PMTK414": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "414"}, // PMTK_API_Q_NMEA_OUTPUT
		"PMTK605": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "605"}, // PMTK_Q_RELEASE
		"PMTK500": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "500"}, // PMTK_DT_FIX_CTL
		"PMTK501": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "501"}, // PMTK_DT_DGPS_MODE
		"PMTK513": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "513"}, // PMTK_DT_SBAS_ENABLED
		"PMTK514": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "514"}, // PMTK_DT_NMEA_OUTPUT
		"PMTK705": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "705"}, // PMTK_DT_RELEASE
		"PMTK869": MtkTypeId{TypeId: TypeId{Talker: TALKER_ID_PROPRIETARY, Code: "MTK"}, PacketType: "869"}, // PMTK_EASY_ENABLE
	}
}

const (
	Valid   DataValid = true
	Invalid DataValid = false
)

type DataValid bool

func (v DataValid) String() string {
	if v == Valid {
		return "A"
	}
	return "V"
}

const (
	NO_FIX                PositioningMode = "N"
	AUTONOMOUS_GNSS_FIX   PositioningMode = "A"
	DIFFERENTIAL_GNSS_FIX PositioningMode = "D"
)

type PositioningMode string

func (p PositioningMode) String() string {
	return string(p)
}

func ParsePositioningMode(raw string) (pm PositioningMode, err error) {
	pm = PositioningMode(raw)
	switch pm {
	case NO_FIX, AUTONOMOUS_GNSS_FIX, DIFFERENTIAL_GNSS_FIX:
	default:
		err = fmt.Errorf("unknow value")
	}
	return
}
