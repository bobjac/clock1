package main

import (
	"io"
	"log"
	"net"
	"time"
	"fmt"
)

func main() {
	fmt.Println("Listening on 0.0.0.0:8000")
	listener, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
