package main

import (
	"log"
	"net"

	"github.com/CaptainFallaway/GoTcpChat/internal/proto"
)

const Version uint8 = 1

type Message struct {
	Username string
	Content  string
}

func handle(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		read, err := conn.Read(buf)
		if err != nil {
			log.Println("Connection closed:", err)
			return
		}

		log.Println("Read", read, "bytes")
		log.Println(err)

		op, err := proto.GetOpCode(buf, Version)
		if err != nil {
			log.Println("Error getting opcode:", err)
			return
		}

		switch op {
		case 1:
			msg, err := proto.Decode[Message](buf)
			if err != nil {
				log.Println("Failed to decode message:", err)
				return
			}
			log.Println(msg)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Recieved connection from", conn.RemoteAddr())
		go handle(conn)
	}
}
