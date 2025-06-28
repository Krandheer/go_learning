package main

import (
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	conn.Read(make([]byte, 1024))
	log.Println("processing the request")
	time.Sleep(10 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, world!\r\n"))
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("listening for new request")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("got new request")

		go do(conn)
	}
}
