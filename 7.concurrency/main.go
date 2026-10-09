package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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

	reader := bufio.NewReader(conn)

	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	line = strings.TrimSuffix(line, "\r\n")
	parts := strings.Split(line, " ")

	if len(parts) != 3 {
		log.Fatal("invalid request line")
	}

	method := parts[0]
	path := parts[1]
	version := parts[2]

	fmt.Printf("----- request -----\n")
	fmt.Printf("raw: %q\n\n", line)

	fmt.Printf("METHOD: %q\n", method)
	fmt.Printf("PATH: %q\n", path)
	fmt.Printf("VERSION: %q\n", version)

	// header parsing
	headers := make(map[string]string)

	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		line = strings.TrimSuffix(line, "\r\n")
		if line == "" {
			break
		}

		parts = strings.SplitN(line, ":", 2)
		key := parts[0]
		value := parts[1]
		headers[key] = value
	}

	fmt.Printf("headers: %q\n", headers)

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
