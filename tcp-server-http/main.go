package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if line == "" {
			// headers are done
			break
		}
		if i == 0 {
			// request line
			reqLineComponents := strings.Fields(line)
			if len(reqLineComponents) == 3 {
				method := reqLineComponents[0]
				uri := reqLineComponents[1]
				mux(conn, method, uri)
			} else {
				break
			}
		}
		i++
	}
}

func mux(conn net.Conn, method string, uri string) {
	statusCode, body := response(method, uri)
	statusLine := responseStatus(statusCode)
	headers := responseHeaders(statusCode, len(body))
	httpResponse := fmt.Sprintf("%s\r\n%s\r\n\n%s", statusLine, headers, body)
	fmt.Fprint(conn, httpResponse)
}

func responseStatus(statusCode int) string {
	var statusLine = ""
	switch statusCode {
	case 200:
		statusLine = "HTTP/1.1 200 OK"
	default:
		statusLine = "HTTP/1.1 404 Not Found"
	}
	return statusLine
}

func responseHeaders(status int, contentLength int) string {
	headers := fmt.Sprintf("Content-Length: %d\r\n", contentLength)
	headers = fmt.Sprintf("%sContent-Type: text/html", headers)
	return headers
}

func response(method string, uri string) (int, string) {
	var statusCode = 200
	var htmlBody = `<p><a href="/">Index</a></p>
<p><a href="/about">About</a></p>
<p><a href="/contact">Contact</a></p>
<p><a href="/apply">Apply</a></p>`

	switch method {
	case "GET":
		switch uri {
		case "/":
			htmlBody = fmt.Sprintf("<h3>INDEX</h3>\r\n%s", htmlBody)
		case "/about":
			htmlBody = fmt.Sprintf("<h3>ABOUT</h3>\r\n%s", htmlBody)
		case "/contact":
			htmlBody = fmt.Sprintf("<h3>CONTACT</h3>\r\n%s", htmlBody)
		case "/apply":
			htmlBody = fmt.Sprintf("<h3>APPLY</h3>\r\n%s", htmlBody)
			applyForm := `<form method="POST" action="/apply">
<input type="submit" value="apply">
</form>`
			htmlBody = fmt.Sprintf("%s\r\n%s", htmlBody, applyForm)
		default:
			statusCode = 404
			htmlBody = "<h2>404 Not Found</h2>"
		}

	case "POST":
		switch uri {
		case "/apply":
			htmlBody = fmt.Sprintf("<h3>APPLY PROCESS</h3>\r\n%s", htmlBody)
		default:
			statusCode = 404
			htmlBody = "<h2>404 Not Found</h2>"
		}

	default:
		statusCode = 404
		htmlBody = "<h2>404 Not Found</h2>"
	}

	htmlStart := `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>TCP Server HTTP MUX</title>
	</head>
	<body>`

	htmlEnd := `</body>
</html>`

	payload := fmt.Sprintf("%s\r\n%s\r\n%s", htmlStart, htmlBody, htmlEnd)

	return statusCode, payload
}
