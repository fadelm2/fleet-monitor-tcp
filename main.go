package main

import (
	"encoding/hex"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	fmt.Println("🚀 TCP Server running on :9000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("🔌 Connected:", conn.RemoteAddr())
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("❌ Disconnected:", conn.RemoteAddr())
			return
		}

		data := buf[:n]
		fmt.Println("📦 RAW:", hex.EncodeToString(data))
		ParseAndLog(data)
	}
}
