# go-nmea [![Go Report Card](https://goreportcard.com/badge/github.com/pilebones/go-nmea)](https://goreportcard.com/report/github.com/pilebones/go-nmea) [![GitHub Repo stars](https://img.shields.io/github/stars/pilebones/go-nmea)](https://somsubhra.github.io/github-release-stats/?username=pilebones&repository=go-nmea) [![GoDoc](https://godoc.org/github.com/pilebones/go-nmea?status.svg)](https://godoc.org/github.com/pilebones/go-nmea) [![Release](https://img.shields.io/github/release/pilebones/go-nmea.svg)](https://github.com/pilebones/go-nmea/releases/latest) [![License](https://img.shields.io/github/license/pilebones/go-nmea)](/LICENSE) ![Build Status](https://github.com/pilebones/go-nmea/workflows/CI/badge.svg) [![Coverage Status](https://coveralls.io/repos/github/pilebones/go-nmea/badge.svg?branch=master)](https://coveralls.io/github/pilebones/go-nmea?branch=master)

A Golang library for decode and serialize standard and proprietary NMEA packet message (GPS information dissector).

Tested with this [GPS Module](http://wiki.52pi.com/index.php/USB-Port-GPS_Module_SKU:EZ-0048) cover [L80 gps protocol specification v1.0.pdf](http://wiki.52pi.com/index.php/File:L80_gps_protocol_specification_v1.0.pdf).
See another [NMEA specification](http://aprs.gids.nl/nmea/).

## NMEA Specification

NMEA standard specification provide 58 kind of message with different structure.
And more according to GPS devices manufacturer (ex: 40 proprietary message identified prefixed by `PMTK` for `L80 GPS protocol specification`).

Syntax: `$<talker_id><message_id>[<data-fields>...]*<checksum><CRLF>`

## Supported NMEA message

The following list will be expanded to manage new types, but now the library can decode and serialize:

* `$GPRMC` - Recommended Minimum Specific GPS/TRANSIT Data
* `$GPVTG` - Track Made Good and Ground Speed
* `$GPGGA` - Global Positioning System Fix Data
* `$GPGSA` - GPS DOP and active satellites
* `$GPGSV` - GPS Satellites in view
* `$GPGLL` - Geographic position, latitude / longitude
* `$GPTXT` - Transfert various text information

## Usage

Library for parsing (read) or serialize (write) NMEA packets (bijective handling), see below:

```go
package main

import "fmt"
import nmea "github.com/pilebones/go-nmea"

func main() {
	raw := "$GPGGA,015540.000,3150.68378,N,11711.93139,E,1,17,0.6,0051.6,M,0.0,M,,*58"

	fmt.Println("Parsing NMEA message:", raw)
	msg, err := nmea.Parse(raw)
	if err != nil {
		fmt.Println("Unable to decode nmea message, err:", err.Error())
		return
	}

	// TODO: Handling complex struct depending on kind of nmea message

	fmt.Println("Craft NMEA packets using Serialize():", msg.Serialize())
}
```

## Documentation
- [GoDoc Reference](http://godoc.org/github.com/pilebones/go-nmea).

## License

go-nmea is available under the [GNU GPL v3 - Clause License](https://opensource.org/licenses/GPL-3.0).
