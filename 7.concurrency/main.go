package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    []byte
}

func parseRequest(reader *bufio.Reader) (Request, error) {
	// === REQUEST INFO ===
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

	headers := make(map[string]string)
	// === PARSING HEADERS ===
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
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		headers[key] = value
	}

	// === READ body ===
	contentLength := 0
	if value, ok := headers["Content-Length"]; ok {
		n, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("invalid Content-Length")
		}

		contentLength = n
	}

	req := Request{
		Method:  method,
		Path:    path,
		Version: version,
		Headers: headers,
	}

	if contentLength > 0 {
		body := make([]byte, contentLength)
		if _, err := io.ReadFull(reader, body); err != nil {
			log.Fatal(err)
		}
		req.Body = body
	}

	return req, nil
}

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
	req, err := parseRequest(reader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("method: %s\n", req.Method)
	fmt.Printf("path: %s\n", req.Path)
	fmt.Printf("version: %s\n", req.Version)
	fmt.Printf("headers: %#v\n", req.Headers)
	fmt.Printf("body: %s\n", req.Body)

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
