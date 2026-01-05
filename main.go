package main

import (
	"encoding/hex"
	"github.com/sirupsen/logrus"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.WithError(err).Fatal("failed to start TCP server")
	}

	log.WithField("port", 9000).
		Info("🚀 TCP Server started")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.WithError(err).Warn("failed to accept connection")
			continue
		}

		log.WithFields(logrus.Fields{
			"remote": conn.RemoteAddr().String(),
		}).Info("🔌 Client connected")

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer func() {
		log.WithField("remote", conn.RemoteAddr().String()).
			Info("❌ Client disconnected")
		conn.Close()
	}()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.WithError(err).
				WithField("remote", conn.RemoteAddr().String()).
				Warn("read error")
			return
		}

		data := buf[:n]

		log.WithFields(logrus.Fields{
			"remote": conn.RemoteAddr().String(),
			"bytes":  n,
			"raw":    hex.EncodeToString(data),
		}).Debug("📦 packet received")

		ParseAndLog(data)
	}
}
