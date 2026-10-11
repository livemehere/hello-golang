package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

type Request struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    []byte
	Params  map[string]string
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

func handleConnection(conn net.Conn, router *Router) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	requestCount := 0

	for {

		// keep alive for 5 sec
		err := conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		// connection close
		if err != nil {
			log.Println("set read deadline:", err)
			return
		}

		req, err := parseRequest(reader)
		if err != nil {
			if err == io.EOF {
				return
			}
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				log.Printf("[TCP idle timeout] %s: 5초 동안 요청이 없어 연결 종료", conn.RemoteAddr())
			} else {
				log.Println("parse request:", err)
			}
			return
		}

		requestCount++
		if requestCount > 1 {
			log.Printf("[TCP reused] %s: 요청 #%d %s %s", conn.RemoteAddr(), requestCount, req.Method, req.Path)
		} else {
			log.Printf("[TCP request] %s: 요청 #%d %s %s", conn.RemoteAddr(), requestCount, req.Method, req.Path)
		}

		res := router.Serve(req)

		if err = writeResponse(conn, res); err != nil {
			log.Println(err)
			return
		}
	}
}

type Router struct {
	routes []Route
}

type Route struct {
	Method  string
	Pattern string
	Handler HandlerFunc
}

type HandlerFunc func(Request) Response

func matchPath(pattern string, path string) (map[string]string, bool) {
	// /users/{id} => ['users', '{id}']
	patternParts := strings.Split(strings.Trim(pattern, "/"), "/")

	// /users/12 => ['users', '12']
	pathParts := strings.Split(strings.Trim(path, "/"), "/")

	if len(patternParts) != len(pathParts) {
		return nil, false
	}

	params := make(map[string]string)

	for i := range patternParts {
		patternPart := patternParts[i]
		pathPart := pathParts[i]

		isParam := strings.HasPrefix(patternPart, "{") && strings.HasSuffix(patternPart, "}")

		if isParam {
			// id
			name := strings.TrimSuffix(strings.TrimPrefix(patternPart, "{"), "}")

			params[name] = pathPart
			continue
		}

		if patternPart != pathPart {
			return nil, false
		}
	}

	return params, true
}

func (r *Router) Handle(method string, pattern string, handler HandlerFunc) {
	r.routes = append(r.routes, Route{
		Method:  method,
		Pattern: pattern,
		Handler: handler,
	})
}

func (r *Router) Serve(req Request) Response {
	for _, route := range r.routes {
		if route.Method != req.Method {
			continue
		}

		params, ok := matchPath(
			route.Pattern,
			req.Path,
		)

		if !ok {
			continue
		}

		req.Params = params

		return route.Handler(req)
	}

	return Response{
		StatusCode: 404,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		Body: []byte("not fount"),
	}
}

func NewRouter() *Router {
	return &Router{
		routes: []Route{},
	}
}

func helloHandler(req Request) Response {
	return Response{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		Body: []byte("hello"),
	}
}

func usersHandler(req Request) Response {
	users := []string{"kong", "Ha"}
	body, err := json.Marshal(users)
	if err != nil {
		return Response{
			StatusCode: 500,
			Body:       []byte("server error"),
		}
	}

	return Response{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		Body: body,
	}
}

func userHandler(req Request) Response {
	id := req.Params["id"]
	return Response{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		Body: []byte(fmt.Sprintf("id is %s", id)),
	}
}

func userAdminHandler(req Request) Response {
	return Response{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		Body: []byte("admin user"),
	}
}

func main() {
	listener, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("listening on :7777")

	router := NewRouter()
	router.Handle("GET", "/hello", helloHandler)
	router.Handle("GET", "/users", usersHandler)
	router.Handle("GET", "/users/{id}", userHandler)
	router.Handle("GET", "/users/admin", userAdminHandler)

	var wg sync.WaitGroup

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		log.Println("[shutdown] shutdown signal received")
		log.Println("[shutdown] stop accepting new connections")

		if err := listener.Close(); err != nil {
			log.Println("[shutdown] listener close:", err)
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		wg.Add(1)

		log.Printf("[TCP new connection] %s", conn.RemoteAddr())
		go func() {
			defer wg.Done()
			handleConnection(conn, router)
		}()
	}
	log.Println("[shubdown] waiting for active connections")
	wg.Wait()
	log.Println("[shutdown] server stopped")
}
