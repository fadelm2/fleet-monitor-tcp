package main

import (
	"encoding/hex"
	"fmt"
)

func ParseAndLog(data []byte) {
	if len(data) < 20 {
		return
	}

	protocol := data[3]

	switch protocol {

	case 0x01:
		imei := parseIMEI(data[4:12])
		fmt.Println("📱 LOGIN")
		fmt.Println("   IMEI:", imei)

	case 0x12:
		lat := parseCoordinate(data[10:14])
		lon := parseCoordinate(data[14:18])

		fmt.Println("📍 GPS DATA")
		fmt.Println("   Latitude :", lat)
		fmt.Println("   Longitude:", lon)
		fmt.Println("   Maps     : https://maps.google.com/?q=",
			lat, ",", lon)
	}
}

func parseCoordinate(b []byte) float64 {
	raw := uint32(b[0])<<24 |
		uint32(b[1])<<16 |
		uint32(b[2])<<8 |
		uint32(b[3])

	return float64(raw) / 30000.0 / 60.0
}

func parseIMEI(b []byte) string {
	return hex.EncodeToString(b)
}
