package main

import (
	"log"
	"net"

	"github.com/CaptainFallaway/GoTcpChat/internal/proto"
)

type Message struct {
	Username string
	Content  string
}

var msg = Message{"CaptainFallaway", "Hello, World!"}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	encoded, _ := proto.Encode(&msg, 1, 1)
	for {
		written, err := conn.Write(encoded)
		log.Println(written, err, encoded)
	}
}
