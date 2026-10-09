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

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       []byte
}

const maxBodySize = 10 * 1024 * 1024

func parseRequest(reader *bufio.Reader) (Request, error) {
	// === REQUEST INFO ===
	line, err := reader.ReadString('\n')
	if err != nil {
		return Request{}, err
	}
	line = strings.TrimSuffix(line, "\r\n")
	parts := strings.Split(line, " ")

	if len(parts) != 3 {
		return Request{}, fmt.Errorf("invalid request line %q", line)
	}

	method := parts[0]
	path := parts[1]
	version := parts[2]

	headers := make(map[string]string)
	// === PARSING HEADERS ===
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			return Request{}, err
		}

		line = strings.TrimSuffix(line, "\r\n")
		if line == "" {
			break
		}

		parts = strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return Request{}, fmt.Errorf("invalid header: %q", line)
		}
		key := strings.ToLower(strings.TrimSpace(parts[0]))
		value := strings.TrimSpace(parts[1])
		headers[key] = value
	}

	// === READ body ===
	contentLength := 0
	if value, ok := headers["content-length"]; ok {
		n, err := strconv.Atoi(value)
		if err != nil {
			return Request{}, err
		}

		if n < 0 {
			return Request{}, fmt.Errorf("invalid Content-Length: %d", n)
		}

		if n > maxBodySize {
			return Request{}, fmt.Errorf("body too large")
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
			return Request{}, err
		}
		req.Body = body
	}

	return req, nil
}

func statusText(code int) string {
	switch code {
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 400:
		return "Bad Request"
	case 404:
		return "Not Found"
	case 500:
		return "Internal Server Error"
	default:
		return "Unknown"
	}
}

func writeResponse(conn net.Conn, res Response) error {
	status := statusText(res.StatusCode)

	// REQUEST LINE
	_, err := fmt.Fprintf(conn, "HTTP/1.1 %d %s\r\n", res.StatusCode, status)
	if err != nil {
		return err
	}

	// HEADERS
	for key, value := range res.Headers {

		// for auto calc
		if strings.EqualFold(key, "Content-Length") {
			continue
		}

		_, err := fmt.Fprintf(conn, "%s: %s\r\n", key, value)
		if err != nil {
			return err
		}
	}
	contentLength := len(res.Body)
	if _, err = fmt.Fprintf(conn, "Content-Length: %d\r\n", contentLength); err != nil {
		return err
	}
	if _, err = fmt.Fprintf(conn, "\r\n"); err != nil {
		return err
	}

	// BODY
	_, err = conn.Write(res.Body)
	if err != nil {
		return err
	}

	return nil
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

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

	res := Response{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		Body: []byte("hello"),
	}

	err = writeResponse(conn, res)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("listening on :7777")

	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("==== client connected:", conn.RemoteAddr().Network(), conn.RemoteAddr().String(), "====")
		go handleConnection(conn)
		fmt.Println("")
	}
}
