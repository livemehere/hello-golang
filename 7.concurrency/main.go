package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("listening on :7777")

	// ---- REQUEST START ----
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("client connected:", conn.RemoteAddr().Network(), conn.RemoteAddr().String())

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("----- request -----\n")
	fmt.Printf("receieved %d bytes\n", n)
	fmt.Println(string(buf[:n]))

	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: 5\r\n" +
		"\r\n" +
		"hello"

	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Fatal(err)
	}

	// ---- REQUEST END ----
}
