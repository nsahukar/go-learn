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

	// write response
	response(conn)
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
			method := strings.Fields(line)[0]
			fmt.Println("METHOD: ", method)
		}
		i++
	}
}

func response(conn net.Conn) {
	body := `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>TCP Server HTTP</title>
	</head>
	<body>
		<h2>Hello, World!</h2>
	</body>
</html>`

	fmt.Fprintln(conn, "HTTP/1.1 200 OK")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintln(conn, "Content-Type: text/html")
	fmt.Fprintln(conn, "")
	fmt.Fprint(conn, body)
}
