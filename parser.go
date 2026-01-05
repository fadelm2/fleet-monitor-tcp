package main

import (
	"encoding/hex"
	"fmt"

	"github.com/sirupsen/logrus"
)

func ParseAndLog(data []byte) {
	if len(data) < 20 {
		log.WithField("length", len(data)).
			Warn("packet ignored (too short)")
		return
	}

	protocol := data[3]

	switch protocol {

	case 0x01:
		imei := parseIMEI(data[4:12])

		log.WithFields(logrus.Fields{
			"protocol": "LOGIN",
			"imei":     imei,
		}).Info("📱 device login")

	case 0x12:
		lat := parseCoordinate(data[10:14])
		lon := parseCoordinate(data[14:18])

		log.WithFields(logrus.Fields{
			"protocol":  "GPS",
			"latitude":  lat,
			"longitude": lon,
			"maps":      "https://maps.google.com/?q=" + formatLatLon(lat, lon),
		}).Info("📍 gps data received")

	default:
		log.WithFields(logrus.Fields{
			"protocol": protocol,
			"raw":      hex.EncodeToString(data),
		}).Warn("unknown protocol")
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

func formatLatLon(lat, lon float64) string {
	return fmt.Sprintf("%f,%f", lat, lon)
}
