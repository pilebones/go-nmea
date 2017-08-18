# go-nmea

A Golang library for decode standard and proprietary NMEA packet message (GPS information dissector).

Tested with this [GPS Module](http://wiki.52pi.com/index.php/USB-Port-GPS_Module_SKU:EZ-0048) cover [L80 gps protocol specification v1.0.pdf](http://wiki.52pi.com/index.php/File:L80_gps_protocol_specification_v1.0.pdf).
See another [NMEA specification](http://aprs.gids.nl/nmea/).

## NMEA Specification

NMEA standard specification provide 58 kind of message with different structure. 
And more according to GPS devices manufacturer (ex: 40 proprietary message identified prefixed by `PMTK` for `L80 GPS protocol specification`).

Syntax: `$<talker_id><message_id>[<data-fields>...]*<checksum><CRLF>`

## Supported NMEA message

_/!\ Work in progress, not a stable release /!\_

The following list will be expanded to decode new types, but now the library can decode only :

* $GPRMC - Recommended Minimum Specific GPS/TRANSIT Data
* $GPVTG - Track Made Good and Ground Speed
* $GPGGA - Global Positioning System Fix Data
* $GPGSA - GPS DOP and active satellites

