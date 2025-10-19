package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {

	ln, _ := net.Listen("tcp", ":6379")

	for {
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("Connection established"))
	fmt.Printf("New connection established from %s to %s\n", conn.RemoteAddr(), conn.LocalAddr())

	store := make(map[string]string)

	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		parts := strings.Fields(strings.TrimSpace(line))

		if len(parts) == 0 {
			continue
		}

		cmd := strings.ToUpper(parts[0])

		switch cmd {
		case "PING":
			conn.Write([]byte("+PONG\r\n"))
		case "SET":
			if len(parts) < 3 {
				conn.Write([]byte(fmt.Sprintf("Wrong number of args, expected: 2, got: %d", len(parts)-1)))
				continue
			}
			key := parts[1]
			value := parts[2]
			store[key] = value
			conn.Write([]byte("Stored"))
		case "GET":
			if len(parts) < 2 {
				conn.Write([]byte(fmt.Sprintf("Wrong number of args, expected: 1, got: %d", len(parts)-1)))
				continue
			}
			key := parts[1]
			value, ok := store[key]
			if !ok {
				conn.Write([]byte("$-1\r\n"))
			} else {
				conn.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(value), value)))
			}
		default:
			conn.Write([]byte("-ERR unknown cmd\r\n"))
		}
	}
}
